package jwt

import (
	"net/http"
	"strings"
)

// ExtractToken Authorization携带token按OAuth2标准带一个bearer
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
