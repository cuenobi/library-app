package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	SecretKey string
	Exp       int
}

func NewJwtToken(cfg Jwt) *Jwt {
	return &Jwt{
		SecretKey: cfg.SecretKey,
		Exp:       cfg.Exp,
	}
}

type CustomClaims struct {
	Username string
	Role     string
	jwt.RegisteredClaims
}

func (j *Jwt) Generate(username, role string) string {
	claims := CustomClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Exp) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return ""
	}

	return ss
}

func (j *Jwt) Validate(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})
}
