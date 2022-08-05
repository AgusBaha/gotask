package models

import (
	"database/sql"
	"fmt"

	"github.com/agusbaha/go-task/config"
	"github.com/agusbaha/go-task/entities"
)

type TaskModel struct {
	conn *sql.DB
}

func NewTaskModel() *TaskModel {
	conn, err := config.DBConnection()

	if err != nil {
		panic(err)
	}

	return &TaskModel{
		conn: conn,
	}
}

func (t *TaskModel) FindAll() ([]entities.Task, error) {
	rows, err := t.conn.Query("SELECT * FROM task")
	if err != nil {
		return []entities.Task{}, err
	}
	defer rows.Close()

	var tasks []entities.Task
	for rows.Next() {
		var task entities.Task
		rows.Scan(&task.Id, &task.TaskDetail, &task.Deadline, &task.Assigment, &task.Status)

		// cara merubah angka di database buat jenis_kelamin/status ke abjad
		// if task.Status == "1" {
		// 	task.Status = "Belum Selesai"
		// } else {
		// 	task.Status = "Selesai"
		// }

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *TaskModel) Create(task entities.Task) bool {
	result, err := t.conn.Exec("INSERT INTO task (task_detail, deadline, assigment, status) VALUES (?, ?, ?, ?)", task.TaskDetail, task.Deadline, task.Assigment, task.Status)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
