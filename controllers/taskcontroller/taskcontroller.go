package taskcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/agusbaha/go-task/entities"
	"github.com/agusbaha/go-task/libraries"
	"github.com/agusbaha/go-task/models"
)

var validation = libraries.NewValidation()
var TaskModel = models.NewTaskModel()

func Index(response http.ResponseWriter, request *http.Request) {
	tasks, _ := TaskModel.FindAll()

	data := map[string]interface{}{
		"Tasks": tasks,
	}

	temp, err := template.ParseFiles("views/task/task.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		// request view jika method get
		temp, err := template.ParseFiles("views/task/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		// mengirim data ke database jika method post
		request.ParseForm()

		var task entities.Task
		task.TaskDetail = request.Form.Get("TaskDetail")
		task.Deadline = request.Form.Get("Deadline")
		task.Assigment = request.Form.Get("Assigment")
		task.Status = request.Form.Get("Status")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(task)

		if vErrors != nil {
			data["task"] = task
			data["validation"] = vErrors
		} else {
			data["Message"] = "Data task berhasil ditambahkan"
			TaskModel.Create(task)
		}

		temp, _ := template.ParseFiles("views/task/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	TaskModel.Delete(id)
	http.Redirect(response, request, "/task", http.StatusSeeOther)
}
