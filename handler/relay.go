package handler

import (
	"fmt"
	"go-relay/common"
	"io"
	"net/http"
)

func RelayHandler(w http.ResponseWriter, r *http.Request) {
	// Check token first.
	if r.Header.Get("X-Relay-Token") != common.CONFIG.Token {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	// Check host.
	if r.Header.Get("X-Relay-Host") == "" {
		http.Error(w, "Invalid host", http.StatusBadRequest)
		return
	}
	host := r.Header.Get("X-Relay-Host")
	protocol := r.Header.Get("X-Relay-Protocol")
	if protocol == "" {
		protocol = "https"
	}
	r.Header.Del("X-Relay-Host")
	r.Header.Del("X-Relay-Token")
	req, err := http.NewRequest(r.Method, fmt.Sprintf("%s://%s%s", protocol, host, r.URL.String()), r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header = r.Header.Clone()
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
