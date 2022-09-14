package main

import (
	"net/http"

	"github.com/DProject89/gocrud/controllers/fotografercontroller"
)

func main() {

	http.HandleFunc("/", fotografercontroller.Index)
	http.HandleFunc("/fotografer", fotografercontroller.Index)
	http.HandleFunc("/fotografer/index", fotografercontroller.Index)
	http.HandleFunc("/fotografer/add", fotografercontroller.Add)
	http.HandleFunc("/fotografer/edit", fotografercontroller.Edit)
	http.HandleFunc("/fotografer/delete", fotografercontroller.Delete)

	http.ListenAndServe(":8989", nil)
}
