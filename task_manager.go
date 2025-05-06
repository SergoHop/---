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

func (tm *TaskManager) CompleteTask(id int) error{
	for _, task := range tm.tasks{
		if task.ID == id{
			task.Complete()
			fmt.Println("задача выполнена")
			return nil
		}
	}
	return fmt.Errorf("такой задачи нету")
}

func (tm *TaskManager) ListTasks(){
	if len(tm.tasks) == 0{
		fmt.Println("задач нету")
		return
	}
	fmt.Print("Список задач")
	for _, task := range tm.tasks{
		fmt.Println(task)
	}
}

func (tm *TaskManager) DeletTask(id int) error{
	for i, task := range tm.tasks{
		if task.ID == id{
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			fmt.Println("задача удалена")
			return nil
		}
	}
	return fmt.Errorf("задача не найдена", id)
}