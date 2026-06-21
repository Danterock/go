package task

import "testing"

func TestFindByIDFoundAndNotFound(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "a", Priority: "low", Done: false},
		{ID: 2, Title: "b", Priority: "high", Done: true},
	}

	found, err := FindByID(tasks, 1)
	if err != nil {
		t.Fatal(err)
	}
	if found.ID != 1 {
		t.Errorf("expected ID 1, got %d", found.ID)
	}

	_, err = FindByID(tasks, 999)
	if err == nil {
		t.Error("expected an error")
	}
}
func TestValidateTitle(t *testing.T) {
	err := ValidateTitle("Go")
	if err != nil {
		t.Error("expected no error")
	}

	err = ValidateTitle("   ")
	if err == nil {
		t.Error("expected an error")
	}
}
