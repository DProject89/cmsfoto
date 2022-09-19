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

	TeamModel := models.Team{}
	teams, err := TeamModel.GetTeams(server.DB)

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

	var team models.Team
	server.DB.First(&team, r.PostForm.Get("id"))
	server.DB.Model(&team).Updates(map[string]interface{}{
		"name":                  r.PostForm.Get("name"),
		"code":                  r.PostForm.Get("code"),
		"location":              r.PostForm.Get("location"),
		"pic_team":              r.PostForm.Get("pic_team"),
		"pic_team_phone_number": r.PostForm.Get("pic_team_phone_number"),
		"pic_team_email":        r.PostForm.Get("pic_team_email"),
		"instagram_team":        r.PostForm.Get("instagram_team"),
		"city":                  r.PostForm.Get("city"),
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

	server.DB.Model(models.Team{}).Create(map[string]interface{}{
		"name":                  r.PostForm.Get("name"),
		"code":                  r.PostForm.Get("code"),
		"location":              r.PostForm.Get("location"),
		"pic_team":              r.PostForm.Get("pic_team"),
		"pic_team_phone_number": r.PostForm.Get("pic_team_phone_number"),
		"pic_team_email":        r.PostForm.Get("pic_team_email"),
		"instagram_team":        r.PostForm.Get("instagram_team"),
		"city":                  r.PostForm.Get("city"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/teams", http.StatusSeeOther)
}

func (server *Server) TeamDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TeamID := vars["id"]

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var team models.Team
	server.DB.Where("id = ?", TeamID).Delete(&team)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/teams", http.StatusSeeOther)
}
