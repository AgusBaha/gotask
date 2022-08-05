package main

import (
	"net/http"

	"github.com/agusbaha/go-task/controllers/taskcontroller"
)

func main() {
	// Route Handler
	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/tasks", taskcontroller.Index)
	http.HandleFunc("/tasks/index", taskcontroller.Index)
	http.HandleFunc("/tasks/add", taskcontroller.Add)
	http.HandleFunc("/tasks/edit", taskcontroller.Edit)
	http.HandleFunc("/tasks/delete", taskcontroller.Delete)

	// Start server
	http.ListenAndServe(":9000", nil)
}
