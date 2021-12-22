package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Authorization(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			err := errors.New("no Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}
		// fmt.Println(r.Header["Token"])

		var mySigningKey = []byte("secretkeyjwt")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error while parsing")
			}
			return mySigningKey, nil
		})

		// fmt.Println(token.Valid)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			err := errors.New("your Token has been expired")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			if claims["role"] == "Admin" {
				r.Header.Set("Role", "Admin")

				handler.ServeHTTP(w, r)
				return
			} else if claims["role"] == "User" {
				r.Header.Set("Role", "User")
				handler.ServeHTTP(w, r)
				return
			}
		}

		err = errors.New("no Authorization")
		json.NewEncoder(w).Encode(err)
	}
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "Admin" {
		w.Write([]byte("Not authorized"))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}
