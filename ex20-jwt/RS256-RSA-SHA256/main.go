package main

import (
	"fmt"
	"log"
	"time"

	"crypto/rand"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// var Secret = []byte("abcdefg")
const Secret = "abcdefg"

func main() {
	// 生成 rsa key对
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatal(fmt.Errorf("generate rsa fail: %w", err))
	}

	tokenString, err := GenerateToken("zhulei", 3600, privateKey)
	if err != nil {
		log.Fatal(fmt.Errorf("generate token fail: %w", err))
	}

	claims, err := ParseToken(tokenString, &privateKey.PublicKey)
	if err != nil {
		log.Fatal(fmt.Errorf("parse token fail: %w", err))
	}

	log.Printf("%#v\n", claims)
}

func GenerateToken(username string, expireSeconds int64, key *rsa.PrivateKey) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		Issuer:    "user service",
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expireSeconds) * time.Second)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        uuid.NewString(),
	},
	)
	s, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}

func ParseToken(tokenString string, key *rsa.PublicKey) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return key, nil
		},
		jwt.WithLeeway(5*time.Second),
	)
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
