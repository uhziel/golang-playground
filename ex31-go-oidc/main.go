package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
)

const issuer = "http://192.168.31.64:8000"
const port = 3111

func randString(bytes int) (string, error) {
	b := make([]byte, bytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, key, value string) {
	cookie := http.Cookie{
		Name:     key,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		panic(err)
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: clientID,
	})

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  fmt.Sprintf("http://192.168.31.64:%d/callback", port),
		Scopes: []string{
			oidc.ScopeOpenID,
			"email",
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		state, err := randString(16)
		if err != nil {
			http.Error(
				w,
				fmt.Errorf("gen state fail: %w", err).Error(),
				http.StatusInternalServerError,
			)
			return
		}

		nonce, err := randString(16)
		if err != nil {
			http.Error(
				w,
				fmt.Errorf("gen nonce fail: %w", err).Error(),
				http.StatusInternalServerError,
			)
			return
		}

		setCallbackCookie(w, r, "state", state)
		setCallbackCookie(w, r, "nonce", nonce)

		authCodeURL := config.AuthCodeURL(
			state,
			oidc.Nonce(nonce),
			oauth2.AccessTypeOffline,
			oauth2.ApprovalForce,
		)
		log.Printf("redirect to %s", authCodeURL)
		http.Redirect(w, r, authCodeURL, http.StatusFound)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		state, err := r.Cookie("state")
		if err != nil {
			http.Error(
				w,
				fmt.Errorf("get cookie state fail: %w", err).Error(),
				http.StatusBadRequest,
			)
			return
		}

		nonce, err := r.Cookie("nonce")
		if err != nil {
			http.Error(
				w,
				fmt.Errorf("get cookie nonce fail: %w", err).Error(),
				http.StatusBadRequest,
			)
			return
		}

		if state.Value != r.URL.Query().Get("state") {
			http.Error(w, "state error", http.StatusInternalServerError)
			return
		}

		code := r.URL.Query().Get("code")

		oauth2token, err := config.Exchange(ctx, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rawIDToken, ok := oauth2token.Extra("id_token").(string) // 通过这么别扭的方式访问是因为 id_token 是 oidc 额外加的
		if !ok {
			http.Error(w, "id_token error", http.StatusInternalServerError)
			return
		}

		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(
				w,
				fmt.Errorf("idToken verify fail: %w", err).Error(),
				http.StatusInternalServerError,
			)
			return
		}

		if idToken.Nonce != nonce.Value {
			http.Error(w, "nonce error", http.StatusInternalServerError)
			return
		}

		/*
			if err := idToken.VerifyAccessToken(oauth2token.AccessToken); err != nil {
				http.Error(w, "accessToken error", http.StatusInternalServerError)
				return
			}
		*/

		resp := struct {
			OAuth2Token *oauth2.Token
			Claims      *json.RawMessage
		}{
			OAuth2Token: oauth2token,
			Claims:      new(json.RawMessage),
		}

		if err := idToken.Claims(&resp.Claims); err != nil {
			http.Error(w, "parse claims fail", http.StatusInternalServerError)
			return
		}

		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			http.Error(w, "MarshalIndent fail"+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	})

	log.Println(fmt.Sprintf("listen at :%d", port))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
