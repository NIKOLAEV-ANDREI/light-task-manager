package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readTaskNumber(scanner *bufio.Scanner) (int, bool) {
	taskNumber, err := strconv.Atoi(readLine(scanner))
	if err != nil {
		fmt.Println("[Введите номер задачи цифрами]")
		return 0, false
	}

	return taskNumber, true
}

type Task struct {
	Title       string
	Description string
	Done        bool
}

func showMenu() {
	fmt.Println()
	fmt.Println("[Менеджер задач]")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Просмотреть задачи")
	fmt.Println("3. Отметить задачу как выполненную")
	fmt.Println("4. Удалить задачу")
	fmt.Println("0. Выйти")
	fmt.Println()

}

func addTask(tasks *[]Task, scanner *bufio.Scanner) {
	var newTask Task

	fmt.Println("[Введите название задачи:]")
	newTask.Title = readLine(scanner)

	fmt.Println("[Введите описание задачи:]")
	newTask.Description = readLine(scanner)

	*tasks = append(*tasks, newTask)
	fmt.Println("[Задача добавлена!]")
	fmt.Println()
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("[Нет задач для отображения!]")
		fmt.Println()
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

func markTaskDone(tasks *[]Task, scanner *bufio.Scanner) {
	if len(*tasks) == 0 {
		fmt.Println("[Нет задач для отметки]")
		return
	}
	showTasks(*tasks)
	fmt.Println("\n[Введите номер задачи для отметки как выполненной:]")
	taskNumber, ok := readTaskNumber(scanner)
	if !ok {
		return
	}
	fmt.Println()

	index := taskNumber - 1

	if index < 0 || index >= len(*tasks) {
		fmt.Println("[Задача с таким номером не найдена!]")
		fmt.Println()
		return
	}

	(*tasks)[index].Done = true
	fmt.Println("[Задача отмечена как выполненная!]")
	fmt.Println()
}

func deleteTask(tasks *[]Task, scanner *bufio.Scanner) {
	if len(*tasks) == 0 {
		fmt.Println("[Нет задач для удаления]")
		return
	}
	showTasks(*tasks)
	fmt.Println("\n[Введите номер задачи для удаления:]")
	taskNumber, ok := readTaskNumber(scanner)
	if !ok {
		return
	}
	fmt.Println()
	index := taskNumber - 1

	if index < 0 || index >= len(*tasks) {
		fmt.Println("[Задача с таким номером не найдена!]")
		fmt.Println()
		return
	}

	fmt.Println("[ВЫ ТОЧНО ХОТИТЕ УДАЛИТЬ ДАННУЮ ЗАДАЧУ?] [Y/N]")
	confirmation := strings.ToLower(readLine(scanner))
	switch confirmation {
	case "n":
		fmt.Println("[Удаление задачи отменено!]")
		fmt.Println()
		return
	case "y":
		*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
		fmt.Println("[Задача удалена!]")
		fmt.Println()
	default:
		fmt.Println("[Неверный ввод. Удаление задачи отменено!]")
		fmt.Println()
		return
	}

}

func main() {
	tasks := []Task{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()
		fmt.Println("[Выберите действие:]")
		choice := readLine(scanner)
		fmt.Println()

		switch choice {
		case "0":
			fmt.Println("[Выход из менеджера задач]")
			return
		case "1":
			addTask(&tasks, scanner)
		case "2":
			showTasks(tasks)
		case "3":
			markTaskDone(&tasks, scanner)
		case "4":
			deleteTask(&tasks, scanner)
		default:
			fmt.Println("[Неизвестная команда. Пожалуйста, выберите действие из меню.]")
			fmt.Println()
		}
	}
}
