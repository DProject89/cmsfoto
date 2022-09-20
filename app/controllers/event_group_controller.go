package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DProject89/cmsfoto/app/models"
	"github.com/gorilla/mux"

	"github.com/unrolled/render"
)

func (server *Server) EventGroups(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	EventGroupModel := models.EventGroup{}
	event_groups, err := EventGroupModel.GetEventGroups(server.DB)

	if err != nil {
		return
	}

	var event_group models.EventGroup
	server.DB.Last(&event_group)

	var groupCode = "GC" + fmt.Sprintf("%03d", event_group.ID+1)

	_ = render.HTML(w, http.StatusOK, "event_groups", map[string]interface{}{
		"event_groups": event_groups,
		"groupCode":    groupCode,
	})
}

func (server *Server) EventGroupEdit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var event_group models.EventGroup
	server.DB.First(&event_group, r.PostForm.Get("id"))
	server.DB.Model(&event_group).Updates(map[string]interface{}{
		"name":           r.PostForm.Get("name"),
		"code":           r.PostForm.Get("code"),
		"location":       r.PostForm.Get("location"),
		"date_start":     r.PostForm.Get("date_start"),
		"date_finish":    r.PostForm.Get("date_finish"),
		"pic_organizer":  r.PostForm.Get("pic_organizer"),
		"organizer_name": r.PostForm.Get("organizer_name"),
		"city":           r.PostForm.Get("city"),
		"event_type":     r.PostForm.Get("event_type"),
		"price":          r.PostForm.Get("price"),
		"event_logo":     r.PostForm.Get("event_logo"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/event-groups", http.StatusSeeOther)
}

func (server *Server) EventGroupAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	server.DB.Model(models.EventGroup{}).Create(map[string]interface{}{
		"name":           r.PostForm.Get("name"),
		"code":           r.PostForm.Get("code"),
		"location":       r.PostForm.Get("location"),
		"date_start":     r.PostForm.Get("date_start"),
		"date_finish":    r.PostForm.Get("date_finish"),
		"pic_organizer":  r.PostForm.Get("pic_organizer"),
		"organizer_name": r.PostForm.Get("organizer_name"),
		"city":           r.PostForm.Get("city"),
		"event_type":     r.PostForm.Get("event_type"),
		"price":          r.PostForm.Get("price"),
		"event_logo":     r.PostForm.Get("event_logo"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/event-groups", http.StatusSeeOther)
}

func (server *Server) EventGroupDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	EventGroupID := vars["id"]

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var event_group models.EventGroup
	server.DB.Where("id = ?", EventGroupID).Delete(&event_group)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/event-groups", http.StatusSeeOther)
}

// func (server *Server) EventGroupCreateEvent(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	EventGroupID := vars["id"]

// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var event models.Event
// 	server.DB.Last(&event)

// 	var eventCode = "E" + fmt.Sprintf("%03d", event.ID+1)

// 	server.DB.Model(models.Event{}).Create(map[string]interface{}{
// 		"event_group_id": EventGroupID,
// 		"code":           eventCode,
// 	})

// 	var event_group models.EventGroup
// 	server.DB.First(&event_group, EventGroupID)
// 	server.DB.Model(&event_group).Updates(map[string]interface{}{
// 		"has_event": 1,
// 	})

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	http.Redirect(w, r, "/events?group="+EventGroupID, http.StatusSeeOther)
// }
