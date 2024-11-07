package storage

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStorage implements TaskStorage for SQLite database storage
type SQLiteStorage struct {
	db *sql.DB
}

// NewSQLiteStorage creates a new SQLite storage
func NewSQLiteStorage() (*SQLiteStorage, error) {
	dbFile, err := getTaskFile("tasks.db")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT FALSE,
        date DATETIME NOT NULL
    )`)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

// SQLite Storage methods
func (ss *SQLiteStorage) AddTask(title string) error {
	currentDate := time.Now()

	_, err := ss.db.Exec("INSERT INTO tasks (title, completed, date) VALUES (?, ?, ?)", title, false, currentDate)
	return err
}

func (ss *SQLiteStorage) ListTasks() ([]Task, error) {
	rows, err := ss.db.Query("SELECT id, title, completed, date FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		// Scan the columns into the Task struct
		if err := rows.Scan(&task.Id, &task.Title, &task.Completed, &task.Date); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ss *SQLiteStorage) DeleteTask(taskId int) error {
	// Yet to be implemented
	return nil
}

func (ss *SQLiteStorage) UpdateTask(taskId int, newTask string) error {
	// Yet to be implemented
	return nil
}

func (ss *SQLiteStorage) TotalTasks() (int, error) {
	var count int
	err := ss.db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	return count, err
}

func (ss *SQLiteStorage) Close() error {
	return ss.db.Close()
}
