package main

import (
	"fmt"
)

type Task struct {
	Title       string
	Description string
	Done        bool
}

func showMenu() {

	fmt.Println("Менеджер задач\n")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Просмотреть задачи")
	fmt.Println("3. Отметить задачу как выполненную")
	fmt.Println("4. Удалить задачу")
	fmt.Println("0. Выйти\n")

}

func main() {
	showMenu()

	var choice string
	fmt.Println("Выберите действие:")
	fmt.Scan(&choice)
	
	switch choice {
	case "0":
		fmt.Println("Выход из менеджера задач")
		return
	case "1":
		fmt.Println("ФУНКЦИЯ-1")
	case "2":
		fmt.Println("ФУНКЦИЯ-2")
	case "3":
		fmt.Println("ФУНКЦИЯ-3")
	case "4":
		fmt.Println("ФУНКЦИЯ-4")
	default:
		fmt.Println("Неизвестная команда")
	}
}
