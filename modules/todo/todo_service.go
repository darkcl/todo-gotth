package todo

import (
	"slices"

	"github.com/pkg/errors"
)

type TodoService struct{
	// Temporary storage for todos
	todos []Todo

	TodoServiceInterface
}

func NewTodoService() TodoServiceInterface {
	return &TodoService{
		todos: []Todo{
			{ID: 1, Title: "Todo 1", Description: "Description 1", Completed: false, CompletedAt: ""},
		},
	}
}

func (s *TodoService) GetTodos() ([]Todo, error) {
	return s.todos, nil
}

func (s *TodoService) GetTodo(id int) (Todo, error) {
	for _, todo := range s.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return Todo{}, errors.Errorf("Todo with ID %d not found", id)
}

func (s *TodoService) CreateTodo(todo Todo) (Todo, error) {
	todo.ID = len(s.todos) + 1
	s.todos = append(s.todos, todo)
	return todo, nil
}

func (s *TodoService) UpdateTodo(id int, todo Todo) (Todo, error) {
	for i, t := range s.todos {
		if t.ID == id {
			t.Title = todo.Title
			t.Description = todo.Description
			t.Completed = todo.Completed
			t.CompletedAt = todo.CompletedAt
			s.todos[i] = t
			return t, nil
		}
	}
	return Todo{}, errors.Errorf("Todo with ID %d not found", id)
}
func (s *TodoService) DeleteTodo(id int) error {
	for i, t := range s.todos {
		if t.ID == id {
			s.todos = slices.Delete(s.todos, i, i+1)
			return nil
		}
	}
	return nil
}
