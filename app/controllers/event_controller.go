package controllers

import (
	"net/http"

	"github.com/DProject89/cmsfoto/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Events(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	teamModel := models.Team{}
	teams, err := teamModel.GetTeams(server.DB)

	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "events", map[string]interface{}{
		"teams": teams,
	})
}
