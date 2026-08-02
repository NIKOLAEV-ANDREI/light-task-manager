package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Title       string
	Description string
	Done        bool
}

func showMenu() {

	fmt.Println("\n[Менеджер задач]")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Просмотреть задачи")
	fmt.Println("3. Отметить задачу как выполненную")
	fmt.Println("4. Удалить задачу")
	fmt.Println("0. Выйти\n")

}

func addTask(tasks *[]Task) {
	var newTask Task
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')

	fmt.Println("[Введите название задачи:]")
	title, _ := reader.ReadString('\n')
	newTask.Title = strings.TrimSpace(title)

	fmt.Println("[Введите описание задачи:]")
	description, _ := reader.ReadString('\n')
	newTask.Description = strings.TrimSpace(description)

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

func markTaskDone(tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println("[Нет задач для отметки]")
		return
	}
	showTasks(*tasks)
	fmt.Println("\n[Введите номер задачи для отметки как выполненной:]")
	var taskNumber int
	fmt.Scan(&taskNumber)
	fmt.Println()

	index := taskNumber - 1

	if index < 0 || index >= len(*tasks) {
		fmt.Println("[Задача с таким номером не найдена!]\n")
		return
	}

	(*tasks)[index].Done = true
	fmt.Println("[Задача отмечена как выполненная!]\n")
}

func deleteTask(tasks *[]Task) {

	if len(*tasks) == 0 {
		fmt.Println("[Нет задач для удаления]")
		return
	}
	showTasks(*tasks)
	fmt.Println("\n[Введите номер задачи для удаления:]")
	var taskNumber int
	fmt.Scan(&taskNumber)
	fmt.Println()
	index := taskNumber - 1

	if index < 0 || index >= len(*tasks) {
		fmt.Println("[Задача с таким номером не найдена!]\n")
		return
	}

	fmt.Println("[ВЫ ТОЧНО ХОТИТЕ УДАЛИТЬ ДАННУЮ ЗАДАЧУ?] [Y/N]")
	var confirmation string
	fmt.Scan(&confirmation)
	switch confirmation {
	case "N", "n":
		fmt.Println("[Удаление задачи отменено!]\n")
		return
	case "Y", "y":
		*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
		fmt.Println("[Задача удалена!]\n")
	default:
		fmt.Println("[Неверный ввод. Удаление задачи отменено!]\n")
		return
	}

}

func main() {
	var choice string
	tasks := []Task{}

	for {
		showMenu()
		fmt.Println("[Выберите действие:]")
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case "0":
			fmt.Println("[Выход из менеджера задач]")
			return
		case "1":
			addTask(&tasks)
		case "2":
			showTasks(tasks)
		case "3":
			markTaskDone(&tasks)
		case "4":
			deleteTask(&tasks)
		default:
			fmt.Println("[Неизвестная команда. Пожалуйста, выберите действие из меню.]\n")
		}
	}
}
