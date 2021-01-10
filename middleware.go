package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func AuthMiddleware(next http.Handler) http.Handler {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Req!")
		// fmt.Println(r.Header["Authorization"])

		authHeaders := r.Header["Authorization"]

		if len(authHeaders) > 0 {
			jwtToken := authHeaders[0]

			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(jwtSecret), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // Set id from recvd token to writer header
				// fmt.Println("Token ID:", claims["id"])
				w.Header().Set("id", fmt.Sprintf("%s", claims["id"]))
			} else {
				fmt.Println(err)
			}
		}

		next.ServeHTTP(w, r)
	})
}

type brotliResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w brotliResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func BrotliMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "br") {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "br")
		br := brotli.NewWriter(w)
		defer br.Close()
		brr := brotliResponseWriter{Writer: br, ResponseWriter: w}
		next.ServeHTTP(brr, r)
	})
}
