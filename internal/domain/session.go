package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/netip"
	"time"
)

// Client errors
var (
	ErrSessionAlreadyLoggedIn = errors.New("already logged in, please logout before logging in")
)

// System errors
var (
	ErrSessionContextNotFound = errors.New("session not found in context")
	ErrSessionDeviceMismatch  = errors.New("device doesn't match with device of current session")
)

type Session struct {
	Id        string    `json:"id"`
	AccountId string    `json:"accountId"`
	UserAgent string    `json:"userAgent"`
	IP        string    `json:"ip"`
	MaxAge    int       `json:"maxAge"`
	ExpiresAt int64     `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewSession(aid, ua, ip string, expiration time.Duration) (*Session, error) {
	// TODO: add useragent validation

	if _, err := netip.ParseAddr(ip); err != nil {
		return nil, fmt.Errorf("netip.ParseAddr: %w", err)
	}

	now := time.Now()

	return &Session{
		AccountId: aid,
		UserAgent: ua,
		IP:        ip,
		MaxAge:    int(expiration.Seconds()),
		ExpiresAt: now.Add(expiration).Unix(),
		CreatedAt: now,
	}, nil
}

func (s *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
