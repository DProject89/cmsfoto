package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/users", server.Users).Methods("GET")
	server.Router.HandleFunc("/user/edit", server.UserEdit).Methods("POST")
	server.Router.HandleFunc("/user/add", server.UserAdd).Methods("POST")
	server.Router.HandleFunc("/user/delete/{id}", server.UserDelete).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
