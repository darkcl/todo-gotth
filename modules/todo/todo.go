package todo

import "github.com/labstack/echo/v4"

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CompletedAt string
}

type TodoControllerInterface interface {
	GetTodos(c echo.Context) error
	GetTodo(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type TodoServiceInterface interface {
	GetTodos() ([]Todo, error)
	GetTodo(id int) (Todo, error)
	CreateTodo(todo Todo) (Todo, error)
	UpdateTodo(id int, todo Todo) (Todo, error)
	DeleteTodo(id int) error
}
