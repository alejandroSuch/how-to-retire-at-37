package main

import (
	"fmt"
	"go-afro/pkg/domain"
)

func main() {
	todos := []domain.ToDo{
		domain.NewTodo("Do the laundry"),
		domain.NewTodo("Learn golang"),
		domain.NewTodo("Have a snack"),
	}

	for i, todo := range todos {
		mark := "✗"

		if i%2 == 1 {
			mark = "✓"
		}

		fmt.Printf("%s %s\n", mark, todo.Name)
	}
}
