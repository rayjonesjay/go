package models

import (
	"context"
	"fmt"
	"log"
	"tasky/db"
	"time"
)

// Task struct defines the schema for the tasks table
type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateTask adds a new task to the database
func CreateTask(userID int, title, description, category string) error {
	query := `
		INSERT INTO tasks (user_id, title, description, category)
		VALUES ($1, $2, $3, $4)
	`
	_, err := db.Conn.Exec(context.Background(), query, userID, title, description, category)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return err
	}
	return nil
}

// GetTasksByUser retrieves all tasks for a specific user
func GetTasksByUser(userID int) ([]Task, error) {
	query := `
		SELECT id, user_id, title, description, is_completed, category, created_at, updated_at
		FROM tasks
		WHERE user_id = $1
	`
	rows, err := db.Conn.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("Error retrieving tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	i := 0
	for rows.Next() {
		i++
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.IsCompleted,
			&task.Category,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning task: %v", err)
			continue
		}
		tasks = append(tasks, task)
	}
	fmt.Println("only", i, "iterations occurred")
	return tasks, nil
}
