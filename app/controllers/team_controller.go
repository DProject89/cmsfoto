package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	var team models.Team
	server.DB.Last(&team)

	var teamCode = "TC" + fmt.Sprintf("%03d", team.ID+1)

	_ = render.HTML(w, http.StatusOK, "teams", map[string]interface{}{
		"teams":    teams,
		"teamCode": teamCode,
	})
}

func (server *Server) TeamEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileExtension := filepath.Ext(handler.Filename)
	filename := r.PostForm.Get("code") + fileExtension
	fileLocation := filepath.Join(dir, "assets/images/uploads", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
		"logo_team":             "public/images/uploads/" + filename,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/teams", http.StatusSeeOther)
}

func (server *Server) TeamAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileExtension := filepath.Ext(handler.Filename)
	filename := r.PostForm.Get("code") + fileExtension
	fileLocation := filepath.Join(dir, "assets/images/uploads", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
		"logo_team":             "assets/images/uploads" + filename,
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
