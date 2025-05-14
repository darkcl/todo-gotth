package routes

import (
	"todo-gotth/modules/todo"

	"github.com/labstack/echo/v4"
)

func SetupAPIRoutes(e *echo.Echo) {
	todo.NewWebTodoRoute(e)
}
