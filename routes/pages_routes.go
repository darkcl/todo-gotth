package routes

import (
	"todo-gotth/modules/todo"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func SetupPagesRoutes(e *echo.Echo) {
	// Setup Main Pages

	e.GET("/", echo.WrapHandler(templ.Handler(todo.Page())))
}
