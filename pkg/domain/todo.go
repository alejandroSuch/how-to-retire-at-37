package domain

import (
	"fmt"
	"time"
)

type ToDo struct {
	Name        string
	completedAt time.Time
}

func (t *ToDo) Complete() {
	t.completedAt = time.Now()
}

func (t ToDo) String() string {
	mark := "✗"
	if !t.completedAt.IsZero() {
		mark = "✓"
	}

	return fmt.Sprintf("%s %s", mark, t.Name)
}

func NewTodo(name string) ToDo {
	return ToDo{Name: name}
}
