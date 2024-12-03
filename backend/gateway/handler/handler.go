package handler

import (
	"github.com/gorilla/mux"
)

func CreateHandler() *Handler {
	mux := mux.NewRouter()
	handler := &Handler{
		Handler: mux,
	}

	mux.HandleFunc("/ping", handler.pingHandler).Methods("GET")
	mux.HandleFunc("/infer/image", handler.imageHandler).Methods("POST")
	mux.HandleFunc("/infer/writer", handler.aiWriterHandler).Methods("POST")
	mux.HandleFunc("/infer/character", handler.characterHandler).Methods("POST")
	mux.HandleFunc("/user", handler.userCheckHandler).Methods("POST")
	mux.HandleFunc("/home", handler.homeHandler).Methods("GET")

	return handler
}
