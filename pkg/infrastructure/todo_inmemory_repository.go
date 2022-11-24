package infrastructure

import (
	"errors"
	"go-afro/pkg/domain"
)

type ToDoInMemoryRepository struct {
	todos []domain.ToDo
}

func NewToDoInMemoryRepository() *ToDoInMemoryRepository {
	return &ToDoInMemoryRepository{}
}

func (r ToDoInMemoryRepository) FindAll() []domain.ToDo {
	return r.todos
}

func (r ToDoInMemoryRepository) FindOne(id int) (domain.ToDo, error) {
	if id < 0 || id > len(r.todos)-1 {
		return domain.ToDo{}, errors.New("element not found")
	}

	return r.todos[id], nil
}

func (r *ToDoInMemoryRepository) Save(t domain.ToDo) error {
	r.todos = append(r.todos, t)

	return nil
}

func (r *ToDoInMemoryRepository) Update(id int, t domain.ToDo) error {
	if id < 0 || id > len(r.todos)-1 {
		return errors.New("element not found")
	}

	r.todos[id] = t

	return nil
}

func (r *ToDoInMemoryRepository) Delete(id int) error {
	if id < 0 || id > len(r.todos)-1 {
		return errors.New("element not found")
	}

	r.todos = append(r.todos[:id], r.todos[id+1:]...)

	return nil
}
