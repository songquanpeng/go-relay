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
		if common.CONFIG.Username == "" || common.CONFIG.Password == "" {
			next.ServeHTTP(w, r)
			return
		}
		// Check username and password
		username, password, ok := r.BasicAuth()
		if !ok || username != common.CONFIG.Username || password != common.CONFIG.Password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
