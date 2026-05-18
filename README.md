# Go CLI Task Manager

A command-line task management application built in Go.

## Features

- Add tasks
- List tasks
- Complete tasks
- Delete tasks
- Persistent JSON storage

## Tech Stack

- Go
- JSON File Storage

## Project Structure

```text
task-manager/
│
├── main.go
├── task/
│   ├── model.go
│   ├── storage.go
│   └── service.go
```

## Run Locally

```bash
go run main.go
```

## Example Commands

```bash
go run main.go add "Learn Go"
go run main.go list
go run main.go complete 1
go run main.go delete 1
