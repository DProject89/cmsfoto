package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DProject89/cmsfoto/app/models"
	"github.com/gorilla/mux"

	"github.com/unrolled/render"
)

func (server *Server) TaggingTypes(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	taggingTypeModel := models.TaggingType{}
	taggingTypes, err := taggingTypeModel.GetTaggingTypes(server.DB)

	if err != nil {
		return
	}

	var taggingType models.TaggingType
	server.DB.Last(&taggingType)

	var taggingTypeCode = "TAG" + fmt.Sprintf("%03d", taggingType.ID+1)

	_ = render.HTML(w, http.StatusOK, "tagging_types", map[string]interface{}{
		"taggingTypes":    taggingTypes,
		"taggingTypeCode": taggingTypeCode,
	})
}

func (server *Server) TaggingTypeEdit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var taggingType models.TaggingType
	server.DB.First(&taggingType, r.PostForm.Get("id"))
	server.DB.Model(&taggingType).Updates(map[string]interface{}{
		"name": r.PostForm.Get("name"),
		"code": r.PostForm.Get("code"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/tagging-types", http.StatusSeeOther)
}

func (server *Server) TaggingTypeAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	server.DB.Model(models.TaggingType{}).Create(map[string]interface{}{
		"name": r.PostForm.Get("name"),
		"code": r.PostForm.Get("code"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/tagging-types", http.StatusSeeOther)
}

func (server *Server) TaggingTypeDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taggingTypeID := vars["id"]

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var taggingType models.TaggingType
	server.DB.Where("id = ?", taggingTypeID).Delete(&taggingType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/tagging-types", http.StatusSeeOther)
}
