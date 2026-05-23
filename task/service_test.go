package task

import (
	"os"
	"testing"
)

func setupTestFile() {
	fileName = "test_tasks.json"
	os.Remove(fileName)
}

func cleanupTestFile() {
	os.Remove(fileName)
}

func TestAddTask(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	err := AddTask("Learn Testing")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	store, err := LoadStore()

	if err != nil {
		t.Fatalf("failed to load store: %v", err)
	}

	if len(store.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(store.Tasks))
	}

	if store.Tasks[0].Title != "Learn Testing" {
		t.Fatalf(
			"expected task title 'Learn Testing', got %s",
			store.Tasks[0].Title,
		)
	}

	if store.Tasks[0].ID != 1 {
		t.Fatalf(
			"expected task ID 1, got %d",
			store.Tasks[0].ID,
		)
	}
}

func TestListTasks(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	AddTask("Task 1")
	AddTask("Task 2")

	tasks, err := ListTasks()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestCompleteTask(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	AddTask("Learn Go")

	err := CompleteTask(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	store, err := LoadStore()

	if err != nil {
		t.Fatalf("failed to load store: %v", err)
	}

	if !store.Tasks[0].Completed {
		t.Fatalf("expected task to be completed")
	}
}

func TestCompleteInvalidTask(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	err := CompleteTask(999)

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestDeleteTask(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	AddTask("Task 1")
	AddTask("Task 2")

	err := DeleteTask(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	store, err := LoadStore()

	if err != nil {
		t.Fatalf("failed to load store: %v", err)
	}

	if len(store.Tasks) != 1 {
		t.Fatalf(
			"expected 1 remaining task, got %d",
			len(store.Tasks),
		)
	}

	if store.Tasks[0].ID != 2 {
		t.Fatalf(
			"expected remaining task ID 2, got %d",
			store.Tasks[0].ID,
		)
	}
}

func TestDeleteInvalidTask(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	err := DeleteTask(999)

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestUpdateTask(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	AddTask("Old Title")

	err := UpdateTask(1, "New Title")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	store, err := LoadStore()

	if err != nil {
		t.Fatalf("failed to load store: %v", err)
	}

	if store.Tasks[0].Title != "New Title" {
		t.Fatalf(
			"expected updated title 'New Title', got %s",
			store.Tasks[0].Title,
		)
	}
}

func TestUpdateInvalidTask(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	err := UpdateTask(999, "Updated")

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestIDDoesNotReuse(t *testing.T) {
    setupTestFile()
    defer cleanupTestFile()

    AddTask("Task 1")  // gets ID 1
    DeleteTask(1)      // deleted
    AddTask("Task 2")  // should get ID 2, NOT 1

    store, _ := LoadStore()

    if store.Tasks[0].ID != 2 {
        t.Fatalf("expected ID 2, got %d — ID was reused", store.Tasks[0].ID)
    }
}

