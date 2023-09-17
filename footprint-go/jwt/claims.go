package jwt

import (
	"github.com/fifpoet/footprint/model"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	User model.UserInfo
	jwt.RegisteredClaims
}
