package api

import (
	"net/http"
	"news/internal/api/http/handlers"
	"news/internal/service"

	"github.com/gorilla/mux"
)

// Создание роутера приложения
func SetupRouter(service service.Usecases) *mux.Router {
	r := mux.NewRouter()
	newsHandler := handlers.NewHandler(service)
	r.Use(HeadersMiddleWare)
	r.HandleFunc("/news/{count:[0-9]+}", newsHandler.NewsHandler).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../../webapp"))))
	return r
}
