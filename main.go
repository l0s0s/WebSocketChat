package main

import (
	"log"
	"net/http"

	"github.com/l0s0s/WebSocketChat/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", &server.ChatHandler{Filename: "index.html"})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("ListenAndServe:", zap.Error(err))
	}
}
