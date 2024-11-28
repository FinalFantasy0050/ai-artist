package main

import (
	"ai-artist/chatbot/handler"
	"ai-artist/chatbot/setting"
	"ai-artist/chatbot/utils/corsController"
	"ai-artist/chatbot/utils/logging"
	"net/http"

	"github.com/urfave/negroni"
)

func initiallization() {
	setting.Init()
	logging.Init()
}

func startServer() {
	mux := handler.CreateHandler()
	handler := negroni.Classic()
	defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	logging.Logger.Info("HTTP server start.")
	http.ListenAndServe(":"+setting.Setting.ServerPort, handler)
}

func main() {
	initiallization()
	startServer()
}
