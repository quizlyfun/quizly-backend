package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"

	"github.com/answersuck/vault/internal/domain/session"

	"github.com/answersuck/vault/pkg/logging"
	"github.com/answersuck/vault/pkg/postgres"
)

const sessionTable = "session"

type sessionRepo struct {
	l logging.Logger
	c *postgres.Client
}

func NewSessionRepo(l logging.Logger, c *postgres.Client) *sessionRepo {
	return &sessionRepo{
		l: l,
		c: c,
	}
}

func (r *sessionRepo) Save(ctx context.Context, s *session.Session) error {
	sql := fmt.Sprintf(`
		INSERT INTO %s (id, account_id, max_age, user_agent, ip, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, sessionTable)

	r.l.Info("psql - session - Save: %s", sql)

	_, err := r.c.Pool.Exec(ctx, sql,
		s.Id,
		s.AccountId,
		s.MaxAge,
		s.UserAgent,
		s.IP,
		s.ExpiresAt,
		s.CreatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return fmt.Errorf("psql - session - Save r.c.Pool.Exec: %w", session.ErrAlreadyExist)
			case pgerrcode.ForeignKeyViolation:
				return fmt.Errorf("psql - session - Save - r.c.Pool.Exec: %w", session.ErrAccountNotFound)
			}
		}

		return fmt.Errorf("psql - session - Save - r.c.Pool.Exec: %w", err)
	}

	return nil
}

func (r *sessionRepo) FindById(ctx context.Context, sessionId string) (*session.Session, error) {
	sql := fmt.Sprintf(`
		SELECT
			s.account_id,
			s.user_agent,
			s.ip,
			s.max_age,
			s.expires_at,
			s.created_at
		FROM %s s
		WHERE s.id = $1
	`, sessionTable)

	r.l.Info("psql - session - FindById: %s", sql)

	s := session.Session{Id: sessionId}

	if err := r.c.Pool.QueryRow(ctx, sql, sessionId).Scan(
		&s.AccountId,
		&s.UserAgent,
		&s.IP,
		&s.MaxAge,
		&s.ExpiresAt,
		&s.CreatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("psql - session - FindById - r.c.Pool.QueryRow.Scan: %w", session.ErrNotFound)
		}

		return nil, fmt.Errorf("psql - session - FindById - r.c.Pool.QueryRow.Scan: %w", err)
	}

	return &s, nil
}

func (r *sessionRepo) FindByIdWithVerified(ctx context.Context, sessionId string) (*session.SessionWithVerified, error) {
	sql := fmt.Sprintf(`
		SELECT
			s.account_id,
			s.user_agent,
			s.ip,
			s.max_age,
			s.expires_at,
			s.created_at,
			a.is_verified
		FROM %s s
		INNER JOIN %s a
		ON s.account_id = a.id
		WHERE s.id = $1
	`, sessionTable, accountTable)

	r.l.Info("psql - session - FindByIdWithVerified: %s", sql)

	s := session.SessionWithVerified{
		Session: session.Session{
			Id: sessionId,
		},
	}

	err := r.c.Pool.QueryRow(ctx, sql, sessionId).Scan(
		&s.Session.AccountId,
		&s.Session.UserAgent,
		&s.Session.IP,
		&s.Session.MaxAge,
		&s.Session.ExpiresAt,
		&s.Session.CreatedAt,
		&s.AccountVerified,
	)
	if err != nil {

		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("psql - session - FindByIdWithVerified - r.c.Pool.QueryRow.Scan: %w", session.ErrNotFound)
		}

		return nil, fmt.Errorf("psql - session - FindByIdWithVerified - r.c.Pool.QueryRow.Scan: %w", err)
	}

	return &s, nil
}

func (r *sessionRepo) FindAll(ctx context.Context, accountId string) ([]*session.Session, error) {
	sql := fmt.Sprintf(`
		SELECT
			id,
			account_id,
			max_age,
			user_agent,
   			ip,
   			expires_at,
			created_at
		FROM %s
		WHERE account_id = $1
	`, sessionTable)

	r.l.Info("psql - session - FindAll: %s", sql)

	rows, err := r.c.Pool.Query(ctx, sql, accountId)
	if err != nil {
		return nil, fmt.Errorf("psql - session - FindAll - r.c.Pool.Query: %w", err)
	}

	defer rows.Close()

	var sessions []*session.Session

	for rows.Next() {
		var s session.Session

		if err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.MaxAge,
			&s.UserAgent,
			&s.IP,
			&s.ExpiresAt,
			&s.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("psql - session - FindAll - rows.Scan: %w", err)
		}

		sessions = append(sessions, &s)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("psql - session - FindAll - rows.Err: %w", err)
	}

	return sessions, nil
}

func (r *sessionRepo) Delete(ctx context.Context, sessionId string) error {
	sql := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, sessionTable)

	r.l.Info("psql - session - Delete: %s", sql)

	ct, err := r.c.Pool.Exec(ctx, sql, sessionId)
	if err != nil {
		return fmt.Errorf("psql - session - Delete - r.c.Pool.Exec: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("psql - session - Delete - r.c.Pool.Exec: %w", session.ErrNotDeleted)
	}

	return nil
}

func (r *sessionRepo) DeleteWithExcept(ctx context.Context, accountId, sessionId string) error {
	sql := fmt.Sprintf(`DELETE FROM %s WHERE account_id = $1 AND id != $2`, sessionTable)

	r.l.Info("psql - session - DeleteWithExcept: %s", sql)

	ct, err := r.c.Pool.Exec(ctx, sql, accountId, sessionId)
	if err != nil {
		return fmt.Errorf("psql - session - Delete - r.c.Pool.Exec: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("psql - session - Delete - r.c.Pool.Exec: %w", session.ErrNotDeleted)
	}

	return nil
}

func (r *sessionRepo) DeleteAll(ctx context.Context, accountId string) error {
	sql := fmt.Sprintf(`DELETE FROM %s WHERE account_id = $1`, sessionTable)

	r.l.Info("psql - session - DeleteAll: %s", sql)

	ct, err := r.c.Pool.Exec(ctx, sql, accountId)
	if err != nil {
		return fmt.Errorf("psql - session - DeleteAll - r.c.Pool.Exec: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("psql - session - DeleteAll - r.c.Pool.Exec: %w", session.ErrNotDeleted)
	}

	return nil
}