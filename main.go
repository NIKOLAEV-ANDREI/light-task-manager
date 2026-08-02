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

	fmt.Println("Менеджер задач")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Просмотреть задачи")
	fmt.Println("3. Отметить задачу как выполненную")
	fmt.Println("4. Удалить задачу")
	fmt.Println("5. Выйти")

}

func main() {
	showMenu()
}
