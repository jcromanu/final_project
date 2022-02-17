package main

import (
	"errors"
	"net/http"
	"strings"

	customErrors "github.com/jcromanu/final_project/http_service/errors"

	"github.com/caarlos0/env/v6"
	"github.com/golang-jwt/jwt"
)

func Authorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, err := GetTokenFromHeader(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		valid, err := ValidateToken(jwt)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		if valid {
			h.ServeHTTP(w, r)
			return
		}
	})
}

func GetTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return "", errors.New(http.StatusText(http.StatusBadRequest))
	} else {
		token = strings.TrimSpace(splitToken[1])
		return token, nil
	}
}

func ValidateToken(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		hmacsecret := &secret{}
		if err := env.Parse(hmacsecret); err != nil {
			return nil, err
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, customErrors.UnexpectedSigningMethod()
		}
		return []byte(hmacsecret.Secret), nil
	})
	if err != nil {
		return false, err
	}
	if token.Valid {
		return true, nil
	}
	return false, nil
}

type secret struct {
	Secret string `env:"HMACSECRET"`
}
