package task

import (
    "errors"
    "time"
)


func AddTask(title string) error {
    store, err := LoadStore()
    if err != nil {
        return err
    }

    newTask := Task{
        ID:        store.NextID,   // use NextID instead of generateTaskID
        Title:     title,
        Completed: false,
        CreatedAt: time.Now(),
    }

    store.Tasks = append(store.Tasks, newTask)
    store.NextID++               

    return SaveStore(store)
}

func ListTasks() ([]Task, error) {
    store, err := LoadStore()
    if err != nil {
        return nil, err
    }
    return store.Tasks, nil
}

func CompleteTask(id int) error {
    store, err := LoadStore()
    if err != nil {
        return err
    }
    for i := range store.Tasks {
        if store.Tasks[i].ID == id {
            store.Tasks[i].Completed = true
            return SaveStore(store)
        }
    }
    return errors.New("task not found")
}

func DeleteTask(id int) error {
    store, err := LoadStore()
    if err != nil {
        return err
    }
    for i := range store.Tasks {
        if store.Tasks[i].ID == id {
            store.Tasks = append(store.Tasks[:i], store.Tasks[i+1:]...)
            return SaveStore(store)    
        }
    }
    return errors.New("task not found")
}

func UpdateTask(id int, newTitle string) error {
    store, err := LoadStore()
    if err != nil {
        return err
    }
    for i := range store.Tasks {
        if store.Tasks[i].ID == id {
            store.Tasks[i].Title = newTitle
            return SaveStore(store)
        }
    }
    return errors.New("task not found")
}
