package main

import "fmt"

type TaskManager struct{
	tasks []*Task
}

//добавляет нью задачу
func (tm *TaskManager) AaddTask(title, description string){
	id := len(tm.tasks) + 1

	task := NewTask(id, title, description)

	tm.tasks = append(tm.tasks, task)
	fmt.Println("добавил задачу")
}