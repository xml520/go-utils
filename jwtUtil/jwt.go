package jwtUtil

import (
	"errors"
	"time"
)
import "github.com/dgrijalva/jwt-go"

type JwtAuth struct {
	ID   int
	Role int
	Ver  int
	Ref  string
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(claims *JwtAuth, secret []byte) string {
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 999999).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// SecretKey 用于对用户数据进行签名，不能暴露
	str, _ := token.SignedString(secret)
	return str
}

// ParseWithClaims 解析Token
func ParseWithClaims(tokenStr string, secret []byte) (*JwtAuth, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtAuth{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtAuth); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("验证不通过")
	}
}
