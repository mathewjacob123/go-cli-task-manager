package task

import "time"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type Store struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}