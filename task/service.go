package task

import "errors"

func generateTaskID(tasks []Task) int {

	maxID := 0

	for _, task := range tasks {

		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}

func AddTask(title string) error {
	tasks, err := LoadTasks()

	if err != nil {
		return err
	}

	newTask := Task{
		ID: generateTaskID(tasks),
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

func UpdateTask(id int, newTitle string) error {

	tasks, err := LoadTasks()

	if err != nil {
		return err
	}

	for i := range tasks {

		if tasks[i].ID == id {

			tasks[i].Title = newTitle

			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}
