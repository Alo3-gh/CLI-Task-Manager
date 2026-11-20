package task

import (
	"fmt"
	"time"
)

const TaskDateFormat = "02-01-2006 15:04:05"

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
	Deadline    time.Time `json:"deadline"`
}

func NewTask(title, description string, deadline time.Time) Task {
	return Task{
		ID:          0,
		Title:       title,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		Deadline:    deadline,
	}
}

func (t *Task) ToString() string {
	return fmt.Sprintf("\nTask: %d\nTitle: %s\nDescription: %s\nStatus: %s\nCreatedAt: %s\nDeadline: %s\n",
		t.ID, t.Title, t.Description, t.Status, t.CreatedAt.Format(TaskDateFormat), t.Deadline.Format(TaskDateFormat))
}
