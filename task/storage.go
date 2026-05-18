package task

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func LoadTasks() ([]Task, error) {
	var tasks []Task

	data, err := os.ReadFile(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	err = json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}