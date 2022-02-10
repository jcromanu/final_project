package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/golang-jwt/jwt"
)

func Authorization(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwt, err := GetTokenFromHeader(r)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			valid, err := ValidateToken(jwt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
			if valid {
				h(w, r)
				return
			}
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
	}
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
			return nil, errors.New("could not parse env variables" + err.Error())
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected method")
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
