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

	server.Router.HandleFunc("/event-groups", server.EventGroups).Methods("GET")
	server.Router.HandleFunc("/event-group/edit", server.EventGroupEdit).Methods("POST")
	server.Router.HandleFunc("/event-group/add", server.EventGroupAdd).Methods("POST")
	server.Router.HandleFunc("/event-group/delete/{id}", server.EventGroupDelete).Methods("GET")

	server.Router.HandleFunc("/events", server.Events).Methods("GET")

	server.Router.HandleFunc("/teams", server.Teams).Methods("GET")
	server.Router.HandleFunc("/team/edit", server.TeamEdit).Methods("POST")
	server.Router.HandleFunc("/team/add", server.TeamAdd).Methods("POST")
	server.Router.HandleFunc("/team/delete/{id}", server.TeamDelete).Methods("GET")

	server.Router.HandleFunc("/roles", server.Roles).Methods("GET")
	server.Router.HandleFunc("/role/edit", server.RoleEdit).Methods("POST")
	server.Router.HandleFunc("/role/add", server.RoleAdd).Methods("POST")
	server.Router.HandleFunc("/role/delete/{id}", server.RoleDelete).Methods("GET")

	server.Router.HandleFunc("/teams", server.Teams).Methods("GET")
	server.Router.HandleFunc("/team/edit", server.TeamEdit).Methods("POST")
	server.Router.HandleFunc("/team/add", server.TeamAdd).Methods("POST")
	server.Router.HandleFunc("/team/delete/{id}", server.TeamDelete).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
