package middleware

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	ErrorServerBusy = "server is busy"
	ErrorReLogin    = "relogin"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	PassWord string `json:"password"`
	UserName string `json:"username"`
}

var (
	Secret     = "123#111" //salt
	ExpireTime = 3600      //token expire time
)

// GenerateJwtToken 生成JwtToken
func GenerateJwtToken(claims *JWTClaims) (string, error) {
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorServerBusy)
	}
	return signedToken, nil
}

// VerifyAction 验证jwt
func VerifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorServerBusy)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReLogin)
	}
	return claims, nil
}
