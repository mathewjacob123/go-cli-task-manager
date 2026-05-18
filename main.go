package main

import (
	"fmt"
	"task-manager/task"
)

func main() {
	tasks, err := task.LoadTasks()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tasks)
}