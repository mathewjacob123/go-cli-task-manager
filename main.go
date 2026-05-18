package main

import (
	"fmt"
	"task-manager/task"
)

func main() {

	err := task.AddTask("Learn Go")

	if err != nil {
		fmt.Println(err)
		return
	}

	tasks, _ := task.ListTasks()

	fmt.Println(tasks)
}