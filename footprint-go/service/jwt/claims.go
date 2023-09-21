package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserId   uint
	UserName string
	jwt.RegisteredClaims
}
