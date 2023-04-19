package handler

import (
	"fmt"
	"go-relay/common"
	"io"
	"net/http"
)

func MirrorHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(r.Method, fmt.Sprintf("%s%s", common.MirrorWebsite, r.URL.String()), r.Body)
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
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
