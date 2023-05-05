package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JWTError = errors.New("error generating JWT")
var InvalidJwtClaims = errors.New("Invalid claims")
var InvalidSigningMethod = errors.New("Invalid signing method, expected HS512")

const oneHrInMs = 3600000

/*
SSOclaims is a struct of JWT claims used for authorization
in E_uprava App.
*/
type SSOclaims struct {
	jwt.RegisteredClaims
}

/*
GenerateJWT generates a json web token (JWT) signed with HS512 method.
Subject is JMBG of logged-in user.
returns the signed token if successful.
JWTError is returned if token couldn't be created
*/
func GenerateJWT(jmbg string, key string) (string, error) {
	claims := SSOclaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(oneHrInMs * time.Millisecond)),
			Subject:   jmbg,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", JWTError
	}
	return tokenString, nil
}

// ParseJwt parses a signed json web token (JWT) string and returns the parsed token
func parseJwt(token string, secret string) (SSOclaims, error) {
	t, err := jwt.ParseWithClaims(token, &SSOclaims{}, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil {
		return SSOclaims{}, err
	}

	if claims, ok := t.Claims.(*SSOclaims); ok {
		return *claims, nil
	}

	return SSOclaims{}, InvalidJwtClaims
}

/*Validate Jwt verifies that jwt is valid
 */
func ValidateJwt(token, secret string) error {
	claims, _ := parseJwt(token, secret)
	if claims.Valid() != nil {
		return InvalidJwtClaims
	}
	return nil
}

func GetPrincipal(token, secret string) (string, error) {
	claims, err := parseJwt(token, secret)
	return claims.Subject, err
}
