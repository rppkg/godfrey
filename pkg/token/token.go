package token

import (
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Options struct {
	SecretKey   string
	IdentityKey string
}

var (
	once sync.Once
	ops  *Options
)

func Init(o *Options) {
	once.Do(func() {
		ops = o
	})
}

func Parse(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(ops.SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	var identity string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		identity = claims[ops.IdentityKey].(string)
	}

	return identity, nil
}

func Sign(identity string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		ops.IdentityKey: identity,
		"nbf":           time.Now().Unix(),
		"iat":           time.Now().Unix(),
		"exp":           time.Now().Add(24 * 7 * time.Hour).Unix(),
	})

	tokenString, err = token.SignedString([]byte(ops.SecretKey))

	return
}
