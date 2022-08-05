package main

import (
	"net/http"

	"github.com/agusbaha/go-task/controllers/taskcontroller"
)

func main() {
	// Route Handler
	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/taks", taskcontroller.Index)
	http.HandleFunc("/taks/index", taskcontroller.Index)
	http.HandleFunc("/task/add", taskcontroller.Add)
	http.HandleFunc("/task/edit", taskcontroller.Edit)
	http.HandleFunc("/task/delete", taskcontroller.Delete)

	// Start server
	http.ListenAndServe(":9000", nil)
}
