package task

import (
	"fmt"
	"time"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusDone     Status = "done"
	StatusArchived Status = "archived"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline,omitempty"`
}

func NewTask(title, description string, deadline time.Time) Task {
	return Task{
		ID:          0, // будет присваиваться в storage
		Title:       title,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		Deadline:    deadline,
	}
}

func (t *Task) ToString() string {
	return fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nStatus: %s\nCreatedAt: %s\nDeadline: %s\n",
		t.ID, t.Title, t.Description, t.Status, t.CreatedAt, t.Deadline)
}
