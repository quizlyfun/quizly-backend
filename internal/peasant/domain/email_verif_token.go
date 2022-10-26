package domain

import (
	"errors"

	"github.com/ysomad/answersuck/jwt"
)

var (
	ErrAccountIDNotFound      = errors.New("account with given id not found")
	ErrEmailVerifTokenExpired = errors.New("email verification token expired")
	ErrEmailAlreadyVerified   = errors.New("email already verified")
)

// EmailVerifToken is a token which must be used to verify user email,
// must be created via constructor only.
type EmailVerifToken string

func NewEmailVerifToken(b jwt.Basic) EmailVerifToken {
	return EmailVerifToken(b)
}

func (t EmailVerifToken) String() string { return string(t) }