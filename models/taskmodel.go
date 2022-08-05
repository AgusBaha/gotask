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

func (t *TaskModel) Create(task entities.Task) bool {
	result, err := t.conn.Exec("INSERT INTO task (task_detail, deadline, assigment, status) VALUES (?, ?, ?, ?)", task.TaskDetail, task.Deadline, task.Assigment, task.Status)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
