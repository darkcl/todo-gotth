package todo

import (
	"net/http"
	// "time"
	"todo-gotth/utils"

	"github.com/labstack/echo/v4"
)

type TodoWebController struct {
	service TodoServiceInterface

	TodoControllerInterface
}

func NewTodoWebController() *TodoWebController {
	return &TodoWebController{
		service: NewTodoService(),
	}
}

func (t *TodoWebController) GetTodos(c echo.Context) error {
	// test: sleep for 2 seconds
	// time.Sleep(2 * time.Second)

	todos, err := t.service.GetTodos()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch todos"})
	}

	return utils.Render(c, http.StatusOK, TodosView(todos))
}

func (t *TodoWebController) GetTodo(c echo.Context) error {
	// Implementation for getting a single todo by ID
	return nil
}

func (t *TodoWebController) CreateTodo(c echo.Context) error {
	// Implementation for creating a new todo
	return nil
}

func (t *TodoWebController) UpdateTodo(c echo.Context) error {
	// Implementation for updating an existing todo
	return nil
}

func (t *TodoWebController) DeleteTodo(c echo.Context) error {
	// Implementation for deleting a todo by ID
	return nil
}
