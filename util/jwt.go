package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/zhanghudong/gopkg/settings"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(username, password string) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(settings.ApplicationConfig.App.Expiration * time.Second)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(settings.ApplicationConfig.App.JwtSecret))
	if err != nil {
		return
	}
	return
}

//解密token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.ApplicationConfig.App.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, err
}
