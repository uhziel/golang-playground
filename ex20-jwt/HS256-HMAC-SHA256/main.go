package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// var Secret = []byte("abcdefg")
const Secret = "abcdefg"

func main() {
	tokenString, err := GenerateToken("zhulei", 3600)
	if err != nil {
		log.Fatal(fmt.Errorf("generate token fail: %w", err))
	}

	claims, err := ParseToken(tokenString)
	if err != nil {
		log.Fatal(fmt.Errorf("parse token fail: %w", err))
	}

	log.Printf("%#v", claims)
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Foo string
}

func GenerateToken(username string, expireSeconds int64) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		jwt.RegisteredClaims{
			Issuer:  "user service",
			Subject: username,
			Audience: jwt.ClaimStrings{
				"app1",
			},
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expireSeconds) * time.Second)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        uuid.NewString(),
		},
		"bar",
	},
	)
	s, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Secret), nil
	},
		jwt.WithIssuer("user service"),
		jwt.WithSubject("zhulei"),
		jwt.WithAudience("app1"),
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(), // 默认不会检查 iss，这里指定后开始检查。检查规则为：当前时间应该在 iss 以后。
		jwt.WithLeeway(3*time.Second),
		jwt.WithValidMethods([]string{"HS256"}),
	)
	// 从实际效果看 token.Valid 用来标识返回的 token 是否有效；如果只是看下 token 是否合法可用，用它就可以。
	// jwt.ParseWithClaims() 返回的 err 可以用来看具体失败的原因是什么。
	// token 可能为 nil，所以这里先判断 err 再使用的 token
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			fmt.Println("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			fmt.Println("Invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		default:
			fmt.Println("Couldn't handle this token:", err)
		}
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("token is not CustomClaims")
	}
	return claims, nil
}
