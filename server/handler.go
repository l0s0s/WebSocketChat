package server

import (
	"net/http"
	"text/template"

	"go.uber.org/zap"
)

type ChatHandler struct {
	Logger *zap.Logger
}

func NewChatHandler(logger *zap.Logger) *ChatHandler {
	return &ChatHandler{
		Logger: logger,
	}
}

func (h *ChatHandler) Index(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		h.Logger.Error("Failed to find template", zap.Error(err))
	}

	err = templ.Execute(w, nil)
	if err != nil {
		h.Logger.Error("Failed to execute all.html", zap.Error(err))
	}
}
