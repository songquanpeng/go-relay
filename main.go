package main

import (
	"fmt"
	"go-relay/common"
	"go-relay/handler"
	"log"
	"net/http"
)

func withMiddleware(middleware func(http.Handler) http.Handler, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware(handler).ServeHTTP(w, r)
	}
}

func main() {
	fmt.Println("Go Relay", common.Version, "is running on port", common.CONFIG.Port)
	mux := http.NewServeMux()
	if common.MirrorMode {
		fmt.Println("Mirror mode is enabled, website:", common.MirrorWebsite)
		mux.HandleFunc("/", withMiddleware(handler.MirrorAuth, handler.MirrorHandler))
	} else {
		mux.HandleFunc("/", withMiddleware(handler.RelayAuth, handler.RelayHandler))
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", common.CONFIG.Port), mux)
	if err != nil {
		log.Fatalln(err)
	}
}
