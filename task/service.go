package task

import "errors"

func AddTask(title string) error {
	tasks, err := LoadTasks()

	if err != nil {
		return err
	}

	newTask := Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, newTask)

	return SaveTasks(tasks)
}

func ListTasks() ([]Task, error) {
	return LoadTasks()
}

func CompleteTask(id int) error {
	tasks, err := LoadTasks()

	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func DeleteTask(id int) error {
	tasks, err := LoadTasks()

	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}
