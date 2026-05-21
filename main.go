package main

import (
	"fmt"
	"os"
	"strconv"

	"task-manager/task"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("task add \"task name\"")
		fmt.Println("task list")
		fmt.Println("task complete <id>")
		fmt.Println("task delete <id>")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":

		if len(os.Args) < 3 {
			fmt.Println("Please provide task title")
			return
		}

		title := os.Args[2]

		err := task.AddTask(title)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task added successfully")

	case "list":

		tasks, err := task.ListTasks()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, t := range tasks {

			status := "❌"

			if t.Completed {
				status = "✅"
			}

			fmt.Printf("%d. %s [%s]\n", t.ID, t.Title, status)
		}

	case "complete":

		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		err = task.CompleteTask(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task completed successfully")

	case "delete":

		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		err = task.DeleteTask(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task deleted successfully")

		case "update":

	if len(os.Args) < 4 {
		fmt.Println("Usage: update <id> <new title>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	newTitle := os.Args[3]

	err = task.UpdateTask(id, newTitle)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task updated successfully")

	default:
		fmt.Println("Unknown command")
	}
}