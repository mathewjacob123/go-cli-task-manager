package task

import (
	"encoding/json"
	"os"
)

var fileName = "tasks.json"

func LoadStore() (Store, error) {
    data, err := os.ReadFile(fileName)
    if err != nil {
        if os.IsNotExist(err) {
            return Store{NextID: 1, Tasks: []Task{}}, nil
        }
        return Store{}, err
    }
    if len(data) == 0 {
        return Store{NextID: 1, Tasks: []Task{}}, nil
    }
    var store Store
    err = json.Unmarshal(data, &store)
    if err != nil {
        return Store{}, err
    }
    return store, nil
}

func SaveStore(store Store) error {
    data, err := json.MarshalIndent(store, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(fileName, data, 0644)
}