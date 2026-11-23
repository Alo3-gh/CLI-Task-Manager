package storage_file

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"task-manager/internal/task"
)

type FileTaskStorage struct {
	filePath string
	mu       sync.Mutex
}

func New(path string) *FileTaskStorage {
	return &FileTaskStorage{filePath: path}
}

func (s *FileTaskStorage) read() ([]task.Task, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}

	var tasks []task.Task
	if len(data) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func (s *FileTaskStorage) write(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

func (s *FileTaskStorage) Create(ctx context.Context, t *task.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks, err := s.read()
	if err != nil {
		return err
	}

	if len(tasks) > 0 {
		t.ID = tasks[len(tasks)-1].ID + 1
	} else {
		t.ID = 1
	}

	tasks = append(tasks, *t)

	return s.write(tasks)
}

func (s *FileTaskStorage) GetAll() ([]task.Task, error) {
	return s.read()
}

func (s *FileTaskStorage) GetByID(taskId int) (task.Task, error) {
	tasks, err := s.read()
	if err != nil {
		return task.Task{}, err
	}

	for _, t := range tasks {
		if t.ID == taskId {
			return t, nil
		}
	}

	return task.Task{}, fmt.Errorf("Task with id: %d not found", taskId)
}
