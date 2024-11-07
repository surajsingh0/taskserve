package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Date      time.Time `json:"-"`
}

// Convert Task to an array of strings
func (t Task) toStringArray() []string {
	return []string{
		fmt.Sprintf("%d", t.Id),
		t.Title,
		fmt.Sprintf("%t", t.Completed),
		t.Date.Format("2006-01-02 15:04:05"),
	}
}

// TaskStorage defines the interface for task storage operations
type TaskStorage interface {
	AddTask(task string, isCompleted bool) error
	ListTasks() ([]Task, error)
	DeleteTask(taskId int) error
	UpdateTask(taskId int, newTask string) error
	ToggleCompleted(taskId int) error
	TotalTasks() (int, error)
	Clear() error
	Close() error
}

func getTaskFile(fileName string) (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	programDir := filepath.Join(configDir, "TaskManager")
	if err := os.MkdirAll(programDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(programDir, fileName), nil
}
