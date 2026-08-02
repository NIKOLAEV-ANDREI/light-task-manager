package main

import (
	"fmt"
)

type Task struct {
	Title       string
	Description string
	Done        bool
}

func main() {
	fmt.Println("Hello, World!")

	tasks := []Task{}
	fmt.Println(tasks)
}
