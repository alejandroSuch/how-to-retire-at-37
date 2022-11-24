package domain

import "time"

type ToDo struct {
	Name        string
	completedAt time.Time
}

func NewTodo(name string) ToDo {
	return ToDo{Name: name}
}
