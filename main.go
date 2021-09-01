package main

import (
	"log"
	"net/http"
	"os"

	"github.com/l0s0s/WebSocketChat/client"
	"github.com/l0s0s/WebSocketChat/server"
	"github.com/l0s0s/WebSocketChat/trace"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	r := client.NewRoom()
	r.Tracer = trace.New(os.Stdout)

	http.Handle("/", &server.ChatHandler{Filename: "chat.html"})
	http.Handle("/room", r)

	go r.Run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("ListenAndServe:", zap.Error(err))
	}
}
