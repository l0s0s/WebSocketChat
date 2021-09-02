package main

import (
	"log"
	"net/http"
	"os"

	"github.com/l0s0s/WebSocketChat/client"
	"github.com/l0s0s/WebSocketChat/server"
	"github.com/l0s0s/WebSocketChat/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	r := client.NewRoom()

	googleO2Config := client.ReadJSON("./client_secret_818196946901-fqevd8f91jpq1e71mfgcnv1e4q6qpcv5.apps.googleusercontent.com.json", logger)
	facebookO2Config := client.ReadJSON("./facebook_conf.json", logger)
	githubO2Config := client.ReadJSON("./github_conf.json", logger)
	r.Tracer = trace.New(os.Stdout)

	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New(googleO2Config["client_id"].(string), googleO2Config["client_secret"].(string),
			googleO2Config["redirect_uris"].([]interface{})[0].(string)),
		facebook.New(facebookO2Config["client_id"].(string), facebookO2Config["client_secret"].(string),
			facebookO2Config["redirect_uris"].([]interface{})[0].(string)),
		github.New(githubO2Config["client_id"].(string), githubO2Config["client_secret"].(string),
			githubO2Config["redirect_uris"].([]interface{})[0].(string)),
	)
	http.Handle("/chat", client.MustAuth(&server.ChatHandler{Filename: "chat.html"}))
	http.Handle("/login", &server.ChatHandler{Filename: "login.html"})
	http.HandleFunc("/auth/", client.LoginHandler)
	http.Handle("/room", r)

	go r.Run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("ListenAndServe:", zap.Error(err))
	}
}
