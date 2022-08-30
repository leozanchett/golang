package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type MyHandler struct{}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Perfect"))
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		authType := strings.Split(auth, " ")
		fmt.Println(authType)
		if len(authType) != 2 || authType[0] != "Basic" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		credentials, err := base64.StdEncoding.DecodeString(authType[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if string(credentials) != "Open Sesame" {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	targetHandler := MyHandler{}
	panic(http.ListenAndServe(":8090", AuthMiddleware(&targetHandler)))
}

// curl localhost:8090 -w "%{http_code}" -H "Authorization: Basic T3BlbiBTZXNhbWU="
// 200
// curl localhost:8090 -w "%{http_code}" - H "Authorization: Basic Hello"
