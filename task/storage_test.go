package task

import (
	"os"
	"testing"
)


func TestSaveAndLoadStore(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	store := Store{
		NextID: 2,
		Tasks: []Task{
			{
				ID:        1,
				Title:     "Test Task",
				Completed: false,
			},
		},
	}

	err := SaveStore(store)

	if err != nil {
		t.Fatalf("failed to save store: %v", err)
	}

	loadedStore, err := LoadStore()

	if err != nil {
		t.Fatalf("failed to load store: %v", err)
	}

	if loadedStore.NextID != 2 {
		t.Fatalf(
			"expected NextID 2, got %d",
			loadedStore.NextID,
		)
	}

	if len(loadedStore.Tasks) != 1 {
		t.Fatalf(
			"expected 1 task, got %d",
			len(loadedStore.Tasks),
		)
	}

	if loadedStore.Tasks[0].Title != "Test Task" {
		t.Fatalf(
			"expected title 'Test Task', got %s",
			loadedStore.Tasks[0].Title,
		)
	}
}

func TestLoadStoreWithCorruptedJSON(t *testing.T) {

	setupTestFile()
	defer cleanupTestFile()

	err := os.WriteFile(
		fileName,
		[]byte("{bad json"),
		0644,
	)

	if err != nil {
		t.Fatalf("failed to write corrupted json")
	}

	_, err = LoadStore()

	if err == nil {
		t.Fatalf("expected JSON parsing error but got nil")
	}
}