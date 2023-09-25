package service

import (
	"errors"
	"github.com/fifpoet/footprint/model"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

const TokenExpireDuration = time.Hour * 2
const TokenRefreshDuration = time.Hour * 24 * 7

// MySecret key must be []byte.
// https://golang-jwt.github.io/jwt/usage/signing_methods/#signing-methods-and-key-types
var MySecret = []byte("tnirptoof") // TODO 生成签名的密钥存入配置中心

type MyClaims struct {
	UserId   uint
	UserName string
	jwt.RegisteredClaims
}

// ExtractToken Authorization携带token按OAuth2标准带一个bearer
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func GenerateToken(userInfo model.User) (string, string, error) {
	expireTime, refreshTime := time.Now().Add(TokenExpireDuration), time.Now().Add(TokenRefreshDuration)
	tokenClaims := &MyClaims{
		UserId:   userInfo.ID,
		UserName: userInfo.Name,
		//字段含义见jwt标准
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "zhaoce",
			Subject:   "login",
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
		},
	}
	refreshClaims := &MyClaims{
		UserId:   userInfo.ID,
		UserName: userInfo.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "zhaoce",
			Subject:   "refresh",
			ExpiresAt: &jwt.NumericDate{Time: refreshTime},
		},
	}
	// 生成Token，指定签名算法和claims
	loginToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	// 签名
	token, err := loginToken.SignedString(MySecret)
	if err != nil {
		return "", "", err
	}
	refresh, err := refreshToken.SignedString(MySecret)
	if err != nil {
		return "", "", err
	}
	return token, refresh, nil
}

// ParseToken 解析JWT到claims中
func ParseToken(tokenString string) (*MyClaims, error) {
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
