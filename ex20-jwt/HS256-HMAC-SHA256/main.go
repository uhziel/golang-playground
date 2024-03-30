package main

import (
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

func GenerateToken(username string, expireSeconds int64) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "user service",
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expireSeconds) * time.Second)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        uuid.NewString(),
	},
	)
	s, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func HS256KeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(Secret), nil
}

func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, HS256KeyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token novalid")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("token is not RegisteredClaims")
	}
	return claims, nil
}
