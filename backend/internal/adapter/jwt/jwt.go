package jwt

import (
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

// Create claims with multiple fields populated
type CustomClaims struct {
	Username string
	jwt.RegisteredClaims
}

// Create a new token object, specifying signing method and the claims
func (j *Jwt) Generate(username string) string {
	claims := CustomClaims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Exp) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(j.SecretKey)
	if err != nil {
		return ""
	}

	return ss
}

func (j *Jwt) Validate(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})
}
