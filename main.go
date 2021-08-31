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
	chatHandler := server.NewChatHandler(logger)
	http.HandleFunc("/", chatHandler.Index)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("ListenAndServe:", zap.Error(err))
	}
}
