package jwt

import (
	"errors"
	"github.com/fifpoet/footprint/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("footprint") // 生成签名的密钥

func GenerateToken(userInfo model.UserInfo) (string, error) {
	expirationTime := time.Now().Add(TokenExpireDuration)
	claims := &MyClaims{
		User: userInfo,
		//TODO  字段含义
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "zhaoce",
			Subject:   "",
			Audience:  nil,
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	}
	// 生成Token，指定签名算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	if tokenString, err := token.SignedString(MySecret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}

}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
