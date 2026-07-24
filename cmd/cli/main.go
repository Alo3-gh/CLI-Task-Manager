package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"task-manager/internal/storage"
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
			//storage.AddTask()
		case "2":
			storage.GetAll()
		// Добавь остальные кейсы: 3 - отметить выполненной, 4 - удалить и т.д.
		case "0":
			fmt.Println("Завершение...")
			return
		}
	}
}

func printMenu() {
	fmt.Println("\n=== Task Manager ===")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Список задач")
	fmt.Println("0. Выход")
}
