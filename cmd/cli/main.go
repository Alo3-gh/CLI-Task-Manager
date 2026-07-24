package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"task-manager/internal/storage"
	"task-manager/internal/task"
)

func main() {
	storage := storage.NewFileStorage("tasks.json")

	for {
		printMenu()
		fmt.Print("Выберите действие: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			addTask(storage)
		case "2":
			showAll(storage)
		case "0":
			fmt.Println("Завершение...")
			return
		}
	}
}

func parseTextInput(scanner *bufio.Scanner, placeholder string) string {
	fmt.Printf("Enter %s: ", placeholder)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func addTask(storage *storage.FileStorage) error {
	scanner := bufio.NewScanner(os.Stdin)

	title := parseTextInput(scanner, "title")
	desc := parseTextInput(scanner, "description")

	deadlineFormat := time.DateOnly
	deadlineStr := parseTextInput(scanner, fmt.Sprintf("deadline date(%s)", deadlineFormat))

	var deadline time.Time
	if deadlineStr != "" {
		d, err := time.Parse(deadlineFormat, deadlineStr)
		if err != nil {
			return err
		}

		deadline = d
	}

	newTask := task.NewTask(title, desc, deadline)
	if err := storage.AddTask(newTask); err != nil {
		return err
	}

	return nil
}

func showAll(storage *storage.FileStorage) error {
	tasks, err := storage.GetAll()
	if err != nil {
		return err
	}

	for _, t := range tasks {
		fmt.Printf("t: %v\n", t.ToString())
	}

	return nil
}

func printMenu() {
	fmt.Println("\n=== Task Manager ===")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Список задач")
	fmt.Println("0. Выход")
}
