package handler

import "github.com/gorilla/mux"

func CreateHandler() *Handler {
	mux := mux.NewRouter()
	handler := &Handler{
		Handler: mux,
	}

	mux.HandleFunc("/ping", handler.pingHandler).Methods("GET")
	mux.HandleFunc("/infer/writer", handler.writerHandler).Methods("POST")
	mux.HandleFunc("/infer/character", handler.characterHandler).Methods("POST")

	return handler
}
