package main

import (
	"net/http"

	"github.com/DProject89/cmsfoto/controllers/fotografercontroller"
)

func main() {
	// Fotografer
	http.HandleFunc("/", fotografercontroller.Index)
	http.HandleFunc("/fotografer", fotografercontroller.Index)
	http.HandleFunc("/fotografer/index", fotografercontroller.Index)
	http.HandleFunc("/fotografer/add", fotografercontroller.Add)
	http.HandleFunc("/fotografer/edit", fotografercontroller.Edit)
	http.HandleFunc("/fotografer/delete", fotografercontroller.Delete)

	// Event
	// http.HandleFunc("/", eventcontroller.Index)
	// http.HandleFunc("/event", eventcontroller.Index)
	// http.HandleFunc("/event/index", eventcontroller.Index)
	// http.HandleFunc("/event/add", eventcontroller.Add)
	// http.HandleFunc("/event/edit", eventcontroller.Edit)
	// http.HandleFunc("/event/delete", eventcontroller.Delete)

	// Event
	// http.HandleFunc("/", eventgroupcontroller.Index)
	// http.HandleFunc("/eventgroup", eventgroupcontroller.Index)
	// http.HandleFunc("/eventgroup/index", eventgroupcontroller.Index)
	// http.HandleFunc("/eventgroup/add", eventgroupcontroller.Add)
	// http.HandleFunc("/eventgroup/edit", eventgroupcontroller.Edit)
	// http.HandleFunc("/eventgroup/delete", eventgroupcontroller.Delete)

	// Event
	// http.HandleFunc("/", evenlistcontroller.Index)
	// http.HandleFunc("/evenlist", evenlistcontroller.Index)
	// http.HandleFunc("/evenlist/index", evenlistcontroller.Index)
	// http.HandleFunc("/evenlist/add", evenlistcontroller.Add)
	// http.HandleFunc("/evenlist/edit", evenlistcontroller.Edit)
	// http.HandleFunc("/evenlist/delete", evenlistcontroller.Delete)

	http.ListenAndServe(":8080", nil)
}
