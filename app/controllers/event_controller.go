package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/DProject89/cmsfoto/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Events(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	eventGroupId := r.URL.Query().Get("group")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	teamModel := models.Team{}
	teams, err := teamModel.GetTeams(server.DB)

	if err != nil {
		return
	}

	eventTeamModel := models.EventTeam{}
	eventTeams, err := eventTeamModel.GetEventTeams(server.DB)

	if err != nil {
		return
	}

	eventTeamByGroup := models.EventTeam{}
	eventTeamsGroup, err := eventTeamByGroup.GetEventTeamsByGroup(server.DB, eventGroupId)
	if err != nil {
		return
	}

	eventModel := models.Event{}
	events, err := eventModel.GetEvents(server.DB)

	if err != nil {
		return
	}

	var event models.Event
	server.DB.Last(&event)

	var event_group models.EventGroup
	server.DB.First(&event_group, eventGroupId)

	var eventCode = "GC" + fmt.Sprintf("%03d", event.ID+1)

	_ = render.HTML(w, http.StatusOK, "events", map[string]interface{}{
		"teams":               teams,
		"events":              events,
		"event_teams":         eventTeams,
		"event_team_by_group": eventTeamsGroup,
		"event_group_id":      eventGroupId,
		"event_code":          eventCode,
		"event_group":         event_group,
	})
}

func (server *Server) EventTeamAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	server.DB.Model(models.EventTeam{}).Create(map[string]interface{}{
		"name":           r.PostForm.Get("team_name"),
		"code":           r.PostForm.Get("team_code"),
		"event_group_id": r.PostForm.Get("event_group_id"),
		"team_id":        r.PostForm.Get("team_id"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/events?group="+r.PostForm.Get("event_group_id"), http.StatusSeeOther)
}

func (server *Server) EventAdd(w http.ResponseWriter, r *http.Request) {
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

	server.DB.Model(models.Event{}).Create(map[string]interface{}{
		"name":            r.PostForm.Get("name"),
		"code":            r.PostForm.Get("code"),
		"description":     r.PostForm.Get("description"),
		"location":        r.PostForm.Get("location"),
		"city":            r.PostForm.Get("city"),
		"date_start":      r.PostForm.Get("date_start"),
		"date_finish":     r.PostForm.Get("date_start"),
		"time_start":      r.PostForm.Get("time_start"),
		"time_finish":     r.PostForm.Get("time_finish"),
		"pic_event":       r.PostForm.Get("pic_event"),
		"event_group_id":  r.PostForm.Get("event_group_id"),
		"team_a_id":       r.PostForm.Get("team_a_id"),
		"team_b_id":       r.PostForm.Get("team_b_id"),
		"payment_method":  r.PostForm.Get("payment_method"),
		"photographer_id": r.PostForm.Get("photographer_id"),
		"price":           r.PostForm.Get("price"),
		"link_youtube":    r.PostForm.Get("link_youtube"),
		"link_media":      r.PostForm.Get("link_media"),
		"cover_image":     "public/images/uploads/" + filename,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/events?group="+r.PostForm.Get("event_group_id"), http.StatusSeeOther)
}
