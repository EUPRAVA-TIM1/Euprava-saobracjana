package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
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

// ParseJwt parses a signed json web token (JWT) string and returns the parsed token
func parseJwt(token string, secret string) (SSOclaims, error) {
	t, err := jwt.ParseWithClaims(token, &SSOclaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, InvalidSigningMethod
		}
		return []byte(secret), nil
	})
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
	claims, err := parseJwt(token, secret)
	if err != nil {
		return err
	}
	
	if claims.Valid() != nil {
		return InvalidJwtClaims
	}

	return nil
}

func GetPrincipal(token, secret string) (string, error) {
	claims, err := parseJwt(token, secret)
	return claims.Subject, err
}
