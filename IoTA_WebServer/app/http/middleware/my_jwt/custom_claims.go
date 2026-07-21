package my_jwt

import "github.com/dgrijalva/jwt-go"

// 自定义jwt的声明字段信息+标准字段
type CustomClaims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}
