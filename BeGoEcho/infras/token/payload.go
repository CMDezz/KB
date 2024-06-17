package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Role      int64     `json:"role"`
	ExpiresAt time.Time `json:"expire_at"`
	IssuedAt  time.Time `json:"created_at"`
	jwt.RegisteredClaims
}

func NewPayload(username string, role int64, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        tokenID,
		Username:  username,
		Role:      role,
		ExpiresAt: time.Now().Add(duration),
		IssuedAt:  time.Now(),
	}, nil

}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return errors.New("token is expired")
	}
	return nil
}
