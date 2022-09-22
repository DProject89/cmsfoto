package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DProject89/cmsfoto/app/models"
	"github.com/gorilla/mux"

	"github.com/unrolled/render"
)

func (server *Server) Taggings(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	taggingModel := models.Tagging{}
	taggings, err := taggingModel.GetTaggings(server.DB)

	if err != nil {
		return
	}

	taggingTypeModel := models.TaggingType{}
	taggingTypes, err := taggingTypeModel.GetTaggingTypes(server.DB)
	if err != nil {
		return
	}

	var tagging models.Tagging
	server.DB.Last(&tagging)

	var taggingCode = "TAG" + fmt.Sprintf("%03d", tagging.ID+1)

	_ = render.HTML(w, http.StatusOK, "taggings", map[string]interface{}{
		"taggings":     taggings,
		"taggingCode":  taggingCode,
		"taggingTypes": taggingTypes,
	})
}

func (server *Server) TaggingEdit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var tagging models.Tagging
	server.DB.First(&tagging, r.PostForm.Get("id"))
	server.DB.Model(&tagging).Updates(map[string]interface{}{
		"name":            r.PostForm.Get("name"),
		"code":            r.PostForm.Get("code"),
		"tagging_type_id": r.PostForm.Get("tagging_type_id"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/taggings", http.StatusSeeOther)
}

func (server *Server) TaggingAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	server.DB.Model(models.Tagging{}).Create(map[string]interface{}{
		"name":            r.PostForm.Get("name"),
		"code":            r.PostForm.Get("code"),
		"tagging_type_id": r.PostForm.Get("tagging_type_id"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/taggings", http.StatusSeeOther)
}

func (server *Server) TaggingDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taggingID := vars["id"]

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var tagging models.Tagging
	server.DB.Where("id = ?", taggingID).Delete(&tagging)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/taggings", http.StatusSeeOther)
}
