package controllers

import (
	"log"
	"net/http"

	"github.com/DProject89/cmsfoto/app/models"
	"github.com/gorilla/mux"

	"github.com/unrolled/render"
)

func (server *Server) Teams(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	teamModel := models.Team{}
	teams, err := teamModel.GetTeams(server.DB)

	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "teams", map[string]interface{}{
		"teams": teams,
	})
}

func (server *Server) TeamEdit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var statusId = 0
	if r.PostForm.Get("status_id") == "on" {
		statusId = 1
	}
	var team models.Team
	server.DB.First(&team, r.PostForm.Get("id"))
	server.DB.Model(&team).Updates(map[string]interface{}{
		"name":      r.PostForm.Get("name"),
		"email":     r.PostForm.Get("email"),
		"status_id": statusId,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/teams", http.StatusSeeOther)
}

func (server *Server) TeamAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var statusId = 0
	if r.PostForm.Get("status_id") == "on" {
		statusId = 1
	}

	server.DB.Model(models.Team{}).Create(map[string]interface{}{
		"name":      r.PostForm.Get("name"),
		"email":     r.PostForm.Get("email"),
		"password":  r.PostForm.Get("password"),
		"status_id": statusId,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/teams", http.StatusSeeOther)
}

func (server *Server) TeamDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["id"]

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var team models.Team
	server.DB.Where("id = ?", teamID).Delete(&team)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/teams", http.StatusSeeOther)
}
