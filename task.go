package main

import (
	"fmt"
	"time"
)

type Task struct{
	ID 			int
	Title 		string
	Description string
	IsCompleted bool
	CreatedAt 	time.Time
}

func NewTask(id int, title, description string) *Task{
	return &Task{
		ID: id,
		Title: title,
		Description: description,
		IsCompleted: false,
		CreatedAt: time.Now(),
	}
}

func (t *Task) Complete(){
	t.IsCompleted = true
}

func (t *Task) String() string{
	status := "не выполнена"
	if t.IsCompleted{
		status = "выполнена"
	}
	return fmt.Sprintf(
	"ID: %d\nЗагаловок: %s\nОписание: %s\nСтатус: %s\nСоздана: %s\n",
	t.ID, t.Title,t.Description, status, t.CreatedAt.Format("2006-01-02 15:00:00"),
	)
}