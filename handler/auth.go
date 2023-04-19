package handler

import (
	"go-relay/common"
	"net/http"
)

func RelayAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check token
		if r.Header.Get("X-Relay-Token") != common.CONFIG.Token {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func MirrorAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: check username and password
		next.ServeHTTP(w, r)
	})
}
