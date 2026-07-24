package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"task-manager/internal/task"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{filePath: path}
}

func (s *FileStorage) Load() ([]task.Task, error) {
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

func (s *FileStorage) Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0644)
}

func (s *FileStorage) AddTask(t task.Task) error {
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	if len(tasks) > 0 {
		t.ID = tasks[len(tasks)-1].ID + 1
	} else {
		t.ID = 1
	}

	tasks = append(tasks, t)

	return s.Save(tasks)
}

func (s *FileStorage) GetAll() ([]task.Task, error) {
	return s.Load()
}

func (s *FileStorage) GetByID(taskId int) (task.Task, error) {
	tasks, err := s.Load()
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
