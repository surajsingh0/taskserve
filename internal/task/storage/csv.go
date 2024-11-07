package storage

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// CSVStorage implements TaskStorage for CSV file storage
type CSVStorage struct {
	file *os.File
}

// NewCSVStorage creates a new CSV storage
func NewCSVStorage() (*CSVStorage, error) {
	taskFile, err := getTaskFile("tasks.csv")
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(taskFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &CSVStorage{file: file}, nil
}

// CSV Storage methods
func (cs *CSVStorage) AddTask(task string) error {
	writer := csv.NewWriter(cs.file)
	defer writer.Flush()

	totalTasks, err := cs.TotalTasks()
	if err != nil {
		log.Println(err)
	}

	newTask := Task{
		Id:        totalTasks,
		Title:     task,
		Completed: false,
		Date:      time.Now(),
	}

	return writer.Write(newTask.toStringArray())
}

func (cs *CSVStorage) ListTasks() ([]Task, error) {
	if _, err := cs.file.Seek(0, 0); err != nil {
		return nil, err
	}

	reader := csv.NewReader(cs.file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, record := range records {
		if len(record) < 4 { // Ensure there are enough fields
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println(err)
		}
		title := record[1]
		completed := record[2] == "true"
		date, _ := time.Parse("2006-01-02 15:04:05", record[3])

		tasks = append(tasks, Task{
			Id:        id,
			Title:     title,
			Completed: completed,
			Date:      date,
		})
	}
	return tasks, nil
}

func (cs *CSVStorage) DeleteTask(taskId int) error {
	tasks, err := cs.ListTasks()
	if err != nil {
		return err
	}

	// Create a temporary file to write the updated tasks
	tempFile, err := os.CreateTemp("", "tasks-*.csv")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name()) // Clean up the temp file after we're done

	writer := csv.NewWriter(tempFile)
	defer writer.Flush()

	taskDeleted := false
	for _, task := range tasks {
		if task.Id != taskId {
			if err := writer.Write(task.toStringArray()); err != nil {
				return err
			}
		} else {
			taskDeleted = true
			fmt.Printf("Deleting task: %+v\n", task)
		}
	}

	if !taskDeleted {
		return fmt.Errorf("task with id %d not found", taskId)
	}

	// Ensure all data is written to the temporary file
	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	if err := tempFile.Close(); err != nil {
		return err
	}

	// Close the original file
	if err := cs.file.Close(); err != nil {
		return err
	}

	// Replace the original file with the temporary file
	if err := os.Rename(tempFile.Name(), cs.file.Name()); err != nil {
		return err
	}

	return nil
}

func (cs *CSVStorage) UpdateTask(taskId int, newTask string) error {
	// Add the new task first
	if err := cs.AddTask(newTask); err != nil {
		return err
	}
	// Delete the specified task
	if err := cs.DeleteTask(taskId); err != nil {
		return err
	}
	return nil
}

func (cs *CSVStorage) TotalTasks() (int, error) {
	if _, err := cs.file.Seek(0, 0); err != nil {
		return 0, err
	}

	reader := csv.NewReader(cs.file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}
	return len(records), nil
}

func (cs *CSVStorage) Close() error {
	return cs.file.Close()
}
