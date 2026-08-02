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

	fmt.Println("[Менеджер задач]\n")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Просмотреть задачи")
	fmt.Println("3. Отметить задачу как выполненную")
	fmt.Println("4. Удалить задачу")
	fmt.Println("0. Выйти\n")

}

func addTask(tasks *[]Task) {
	var newTask Task

	fmt.Println("[Введите название задачи:]")
	fmt.Scan(&newTask.Title)
	fmt.Println("[Введите описание задачи:]")
	fmt.Scan(&newTask.Description)

	*tasks = append(*tasks, newTask)
	fmt.Println("[Задача добавлена!]\n")
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("[Нет задач для отображения!]\n")
	} else {
		fmt.Println("[Список задач]:")
		for i, task := range tasks {
			status := "Не выполнена"
			if task.Done {
				status = "Выполнена"
			}
			fmt.Printf("%d. %s - %s [%s]\n", i+1, task.Title, task.Description, status)
		}
	}
}

func main() {
	var choice string
	tasks := []Task{}

	for {
		showMenu()
		fmt.Println("[Выберите действие:]")
		fmt.Scan(&choice)

		switch choice {
		case "0":
			fmt.Println("[Выход из менеджера задач]")
			return
		case "1":
			addTask(&tasks)
		case "2":
			showTasks(tasks)
		case "3":
			fmt.Println("[ФУНКЦИЯ-3]")
		case "4":
			fmt.Println("[ФУНКЦИЯ-4]")
		default:
			fmt.Println("[Неизвестная команда. Пожалуйста, выберите действие из меню.]\n")
		}
	}
}
