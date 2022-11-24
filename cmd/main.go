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

	todos[1].Complete()

	for _, todo := range todos {
		fmt.Printf("%s\n", todo)
	}
}
