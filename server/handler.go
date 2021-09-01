package server

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"go.uber.org/zap"
)

// ChatHandler represents a single template.
type ChatHandler struct {
	Logger   *zap.Logger
	Once     sync.Once
	Filename string
	Templ    *template.Template
}

func (h *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Once.Do(func() {
		h.Templ = template.Must(template.ParseFiles(filepath.Join("templates",
			h.Filename)))
	})
	h.Templ.Execute(w, r)
}
