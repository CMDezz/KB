package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/CMDezz/KB/utils/constants"
	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenMaker struct {
	secretKey string
}

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

func NewJWTTokenMaker(secretKey string) (*JWTTokenMaker, error) {
	if len(secretKey) < constants.MinLengthSecretKey {
		return nil, fmt.Errorf("invalid keysize: secret key must be atleast %d characters", constants.MinLengthSecretKey)
	}
	return &JWTTokenMaker{
		secretKey: secretKey,
	}, nil
}

func (token *JWTTokenMaker) NewToken(username string, role int64, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", &Payload{}, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, *payload)
	tk, err := jwtToken.SignedString([]byte(token.secretKey))

	if err != nil {
		return "", &Payload{}, err
	}
	return tk, payload, nil
}

func (token *JWTTokenMaker) ValidToken(tk string) (*Payload, error) {
	var keyFunc = func(tk *jwt.Token) (any, error) {
		//token khong phai method security da chon
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(token.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(tk, &Payload{}, keyFunc)
	if err != nil {
		// verr, ok := err.(*jwt.)
		// if ok && errors.Is(verr.Inner, ErrExpiredToken) {
		// 	return nil, ErrExpiredToken
		// }
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil

}
