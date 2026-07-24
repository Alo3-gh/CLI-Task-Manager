package storage

import (
	"encoding/json"
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

func (s *FileStorage) AddTask(task task.Task) error {

	return nil
}

func (s *FileStorage) GetAll() error {
	return nil
}

func (s *FileStorage) GetByID(taskId int) error {
	return nil
}

func (s *FileStorage) Update(taskId int) error {
	return nil
}

func (s *FileStorage) Delete(taskId int) error {
	return nil
}

// Добавить методы: Add, GetAll, GetByID, Update, Delete
