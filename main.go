package main

import (
	"log"
	"net/http"
	"os"

	qrgo "github.com/elsesiy/qrgo/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	handler := http.HandlerFunc(qrgo.QRServer)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("couldn't bind port %s - %v", port, err)
	}
}
