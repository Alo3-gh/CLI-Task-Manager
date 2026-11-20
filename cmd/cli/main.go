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

func printMenu() {
	fmt.Println("\n=== Task Manager ===")
	fmt.Println("1. Add task")
	fmt.Println("2. Show tasks")
	fmt.Println("0. Exit")
}

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

func scanTextInput(scanner *bufio.Scanner, placeholder string) string {
	fmt.Printf("Enter %s: ", placeholder)
	scanner.Scan()

	return strings.TrimSpace(scanner.Text())
}

func addTask(storage *storage.FileStorage) error {
	scanner := bufio.NewScanner(os.Stdin)

	title := scanTextInput(scanner, "title")
	desc := scanTextInput(scanner, "description")
	deadlineStr := scanTextInput(scanner, fmt.Sprintf("deadline date(%s)", task.TaskDateFormat))

	deadline, err := time.Parse(task.TaskDateFormat, deadlineStr)
	if err != nil {
		deadline = time.Now().Add(time.Hour * 24)
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
		fmt.Printf("%v\n", t.ToString())
	}

	return nil
}
