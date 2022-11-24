package domain

type TodoRepository interface {
	FindAll() []ToDo
	FindOne(id int) (ToDo, error)
	Save(ToDo) error
	Update(int, ToDo) error
	Delete(int) error
}
