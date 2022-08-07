package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/agusbaha/go-task/config"
	"github.com/agusbaha/go-task/entities"
)

type TaskModel struct {
	conn *sql.DB
}

// index task
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
		rows.Scan(&task.Id, &task.TaskDetail, &task.Assigment, &task.Deadline, &task.Status)

		// cara merubah angka di database buat jenis_kelamin/status ke abjad
		// if task.Status == "1" {
		// 	task.Status = "Belum Selesai"
		// } else {
		// 	task.Status = "Selesai"
		// }

		// Format date status
		// 2005-01-01 yyyy-mm-dd
		Deadline, _ := time.Parse("2006-01-02", task.Deadline)
		// 01-01-2005 dd-mm-yyyy
		task.Deadline = Deadline.Format("02-01-2006")

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Create a new task
func (t *TaskModel) Create(task entities.Task) bool {
	result, err := t.conn.Exec("INSERT INTO task (task_detail, deadline, assigment, status) VALUES (?, ?, ?, ?)", task.TaskDetail, task.Deadline, task.Assigment, task.Status)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

// Find a task by id
func (t *TaskModel) Find(id int64, task *entities.Task) error {
	return t.conn.QueryRow("SELECT * FROM task WHERE id = ?", id).Scan(&task.Id, &task.TaskDetail, &task.Assigment, &task.Deadline, &task.Status)
}

// Update a task by id
func (t *TaskModel) Update(task entities.Task) error {
	_, err := t.conn.Exec("UPDATE task SET task_detail = ?, deadline = ?, assigment = ?, status = ? WHERE id = ?", task.TaskDetail, task.Deadline, task.Assigment, task.Status, task.Id)

	if err != nil {
		return err
	}

	return nil
}
func (t *TaskModel) Delete(id int64) {
	t.conn.Exec("DELETE FROM task WHERE id = ?", id)
}
