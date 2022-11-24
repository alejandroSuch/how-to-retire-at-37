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

func (t ToDo) IsCompleted() bool {
	return !t.completedAt.IsZero()
}

func (t ToDo) String() string {
	mark := "✗"
	if t.IsCompleted() {
		mark = "✓"
	}

	return fmt.Sprintf("%s %s", mark, t.Name)
}

func NewTodo(name string) ToDo {
	return ToDo{Name: name}
}
