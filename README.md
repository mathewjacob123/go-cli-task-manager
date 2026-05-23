# Go CLI Task Manager

A command-line task management application built in Go that lets you manage tasks from your terminal with persistent JSON storage.

## Features

- Add tasks with automatic ID assignment
- List all tasks with completion status
- Complete, update, and delete tasks
- Persistent storage — tasks survive between sessions
- IDs never reuse even after deletion

## Tech Stack

- Go 1.21+
- JSON file storage (no database required)

## Project Structure

text
task-manager/
├── main.go           # Entry point, CLI command routing
└── task/
    ├── model.go      # Task and Store struct definitions
    ├── storage.go    # JSON read/write operations
    └── service.go    # Business logic (add, list, complete, delete, update)
```



## Usage

bash
# Add a task
go run main.go add "Learn Go"

# List all tasks
go run main.go list

# Complete a task
go run main.go complete 1

# Update a task
go run main.go update 1 "Learn Go and build REST APIs"

# Delete a task
go run main.go delete 1


## Example Output


1. Learn Go [❌]
2. Build REST API [✅]

## Installation
git clone https://github.com/mathewjacob123/go-cli-task-manager
cd go-cli-task-manager
go run main.go


## What I Learned

- Go project structure and package organisation
- Structs, interfaces, and error handling in Go
- File I/O and JSON marshalling/unmarshalling
- Slice manipulation patterns in Go
- Writing unit tests with Go's testing package
- Persistent ID management to prevent ID reuse