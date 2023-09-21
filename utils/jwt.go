package utils

import (
	"fmt"
	"reggie/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(data interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data":  data,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * time.Duration(global.ConfigViper.GetInt("token.time"))).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(global.ConfigViper.GetString("token.secret")))
	return tokenString
}

// 解析token的函数
func ParseToken(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.ConfigViper.GetString("token.secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	data := claims["data"].(map[string]any)
	return data, err
}
