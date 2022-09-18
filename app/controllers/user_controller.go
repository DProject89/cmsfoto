package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DProject89/cmsfoto/app/models"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"github.com/unrolled/render"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.initializeDB(dbConfig)
	server.initializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("Listening to port %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host= %s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database server")
	}
}

func (server *Server) Users(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to fotogrit")
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})

	userModel := models.User{}
	users, err := userModel.GetUsers(server.DB)

	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "users", map[string]interface{}{
		"users": users,
	})
}

func (server *Server) UserEdit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var statusId = 0
	if r.PostForm.Get("status_id") == "on" {
		statusId = 1
	}
	var user models.User
	server.DB.First(&user, r.PostForm.Get("id"))
	server.DB.Model(&user).Create(map[string]interface{}{
		"name":      r.PostForm.Get("name"),
		"email":     r.PostForm.Get("email"),
		"status_id": statusId,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (server *Server) UserAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var statusId = 0
	if r.PostForm.Get("status_id") == "on" {
		statusId = 1
	}

	server.DB.Model(models.User{}).Create(map[string]interface{}{
		"id":        uuid.NewV4(),
		"name":      r.PostForm.Get("name"),
		"email":     r.PostForm.Get("email"),
		"password":  r.PostForm.Get("password"),
		"status_id": statusId,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (server *Server) UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	server.DB.Where("id = ?", userID).Delete(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
