package task

import (
	"fmt"
	"todo/internal/storage_type"
	"todo/internal/task/storage"
)

// TaskManager handles task operations using the chosen storage
type TaskManager struct {
	storage storage.TaskStorage
}

// NewTaskManager creates a new TaskManager with the specified storage type
func NewTaskManager(storageType storage_type.StorageType) (*TaskManager, error) {
	var taskStorage storage.TaskStorage
	var err error

	switch storageType {
	case storage_type.CSV:
		taskStorage, err = storage.NewCSVStorage()
	case storage_type.SQLite:
		taskStorage, err = storage.NewSQLiteStorage()
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", storageType)
	}

	if err != nil {
		return nil, err
	}

	return &TaskManager{storage: taskStorage}, nil
}

// AddTask adds a new task
func (tm *TaskManager) AddTask(task string) error {
	return tm.storage.AddTask(task)
}

// ListTasks lists all tasks
func (tm *TaskManager) ListTasks() ([]storage.Task, error) {
	return tm.storage.ListTasks()
}

// Delete a task
func (tm *TaskManager) DeleteTask(taskId int) error {
	return tm.storage.DeleteTask(taskId)
}

// Update a task
func (tm *TaskManager) UpdateTask(taskId int, newTask string) error {
	return tm.storage.UpdateTask(taskId, newTask)
}

// TotalTasks returns the total number of tasks
func (tm *TaskManager) TotalTasks() (int, error) {
	return tm.storage.TotalTasks()
}

// Close closes the storage
func (tm *TaskManager) Close() error {
	return tm.storage.Close()
}
