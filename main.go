package main

import (
	"fmt"
	"go-relay/common"
	"go-relay/handler"
	"log"
	"net/http"
)

func main() {
	// Check common/init.go for initialization.
	fmt.Println("Go Relay", common.Version, "is running on port", common.CONFIG.Port)
	http.HandleFunc("/", handler.RelayHandler)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", common.CONFIG.Port), nil))
}
