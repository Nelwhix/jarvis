package todo_test

import (
	"testing"
	"github.com/Nelwhix/todo"
	"os"
)

func TestAdd(t *testing.T) {
	list := todo.List{}

	taskName := "Test task"
	list.Add(taskName)

	if list[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, list[0].Task)
	}
}

func TestComplete(t *testing.T) {
	list := todo.List{}
	taskName := "Test Task"
	list.Add(taskName)

	if list[0].Done {
		t.Errorf("New task should not be completed.")
	}
	list.Complete(1)

	if !list[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	list := todo.List{}

	tasks := []string{
		"Test task 1",
		"Test task 2",
		"Test task 3",
	}

	for _, v := range tasks {
		list.Add(v)
	}

	list.Delete(2)

	if len(list) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(list))
	}

	if list[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], list[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	list1 := todo.List{}
	list2 := todo.List{}

	taskName := "Test task"
	list1.Add(taskName)

	tempFile, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tempFile.Name())

	if err := list1.Save(tempFile.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := list2.Get(tempFile.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if list1[0].Task != list2[0].Task {
		t.Errorf("Task %q should match %q task.", list1[0].Task, list2[0].Task)
	}
}