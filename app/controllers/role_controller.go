package controllers

import (
	"log"
	"net/http"

	"github.com/DProject89/cmsfoto/app/models"
	"github.com/gorilla/mux"

	"github.com/unrolled/render"
)

func (server *Server) Roles(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	roleModel := models.Role{}
	roles, err := roleModel.GetRoles(server.DB)

	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "roles", map[string]interface{}{
		"roles": roles,
	})
}

func (server *Server) RoleEdit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var statusRole = 0
	if r.PostForm.Get("status_role") == "on" {
		statusRole = 1
	}

	var role models.Role
	server.DB.First(&role, r.PostForm.Get("id"))
	server.DB.Model(&role).Updates(map[string]interface{}{
		"role":        r.PostForm.Get("role"),
		"description": r.PostForm.Get("description"),
		"status_role": statusRole,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/roles", http.StatusSeeOther)
}

func (server *Server) RoleAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var statusRole = 0
	if r.PostForm.Get("status_role") == "on" {
		statusRole = 1
	}

	server.DB.Model(models.Role{}).Create(map[string]interface{}{
		/* "id":          uuid.NewV4(), */
		"role":        r.PostForm.Get("role"),
		"description": r.PostForm.Get("description"),
		"status_role": statusRole,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/roles", http.StatusSeeOther)
}

func (server *Server) RoleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var role models.Role
	server.DB.Where("id = ?", roleID).Delete(&role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/roles", http.StatusSeeOther)
}
