package storage

import (
	"os"
	"path/filepath"
	"testing"
	"week3-tasker/task"
)

func TestLoadMissingFileReturnsEmptySlice(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing.json")
	tasks, err := LoadTasksJSON(path)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(tasks) != 0 {
		t.Fatalf("expected empty slice, got %d tasks", len(tasks))
	}
}
func TestLoadEmptyFileReturnsEmptySlice(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.json")
	err := os.WriteFile(path, []byte(""), 0644)
	if err != nil {
		t.Fatal(err)
	}
	tasks, err := LoadTasksJSON(path)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(tasks) != 0 {
		t.Fatalf("expected empty slice, got %d tasks", len(tasks))
	}
}
func TestLoadBrokenJSONReturnsError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "broken.json")
	err := os.WriteFile(path, []byte("{bad JSON"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	_, err = LoadTasksJSON(path)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

}
func TestSaveAndLoadRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.json")
	tasks := []task.Task{
		{
			ID:       1,
			Title:    "a",
			Priority: "low",
			Done:     false,
		},
		{
			ID:       2,
			Title:    "b",
			Priority: "medium",
			Done:     true,
		},
		{
			ID:       3,
			Title:    "c",
			Priority: "high",
			Done:     false,
		},
	}
	err := SaveTasksJSON(path, tasks)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	expected, err := LoadTasksJSON(path)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(expected) != len(tasks) {
		t.Fatalf("expected %d tasks, got %d", len(tasks), len(expected))
	}
	if expected[0].Title != tasks[0].Title {
		t.Fatalf("expected %s, got %s",
			tasks[0].Title,
			expected[0].Title)
	}
	if expected[1].Done != tasks[1].Done {
		t.Fatalf("expected %t, got %t",
			tasks[1].Done,
			expected[1].Done)
	}
}
