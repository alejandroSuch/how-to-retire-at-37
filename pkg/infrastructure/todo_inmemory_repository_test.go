package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"go-afro/pkg/domain"
	"testing"
)

func TestToDoInMemoryRepository_Update_Ok(t *testing.T) {
	// ARRANGE
	r := ToDoInMemoryRepository{
		todos: []domain.ToDo{
			domain.NewTodo("todo1"),
			domain.NewTodo("todo2"),
			domain.NewTodo("todo3"),
		},
	}

	todo := domain.NewTodo("another")
	todo.Complete()

	// ACT
	err := r.Update(1, todo)

	// ASSERT
	assert.Nil(t, err)
	assert.True(t, r.todos[1].IsCompleted())
	assert.Equal(t, "another", r.todos[1].Name)
}

func TestToDoInMemoryRepository_Delete_Ok(t *testing.T) {
	// ARRANGE
	r := ToDoInMemoryRepository{
		todos: []domain.ToDo{
			domain.NewTodo("todo1"),
			domain.NewTodo("todo2"),
			domain.NewTodo("todo3"),
		},
	}

	// ACT
	err := r.Delete(2)

	// ASSERT
	assert.Nil(t, err)
	assert.Equal(t, 2, len(r.todos))
	assert.Equal(t, "todo1", r.todos[0].Name)
	assert.Equal(t, "todo2", r.todos[1].Name)

}
