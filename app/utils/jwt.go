package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenUser 签名需要传递的参数(根据自己定义)
type TokenUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type MyClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// LoginStruct 登录的参数
type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 证书签名密钥
var jwtKey = []byte("f562c8543707aa0d")

// GenerateToken 定义生成token的方法
func GenerateToken(u TokenUser) (string, error) {
	// 定义过期时间,7天后过期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &MyClaims{
		UserId:   u.Id,
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 发布时间
			Subject:   "token",               // 主题
			Issuer:    "dcr-gin",             // 发布者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 定义解析token的方法
func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
