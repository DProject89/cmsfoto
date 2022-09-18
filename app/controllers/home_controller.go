package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "Home Title",
		"body":  "Home Description",
	})
}
