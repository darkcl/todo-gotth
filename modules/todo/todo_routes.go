package todo

import (
	"github.com/labstack/echo/v4"
)

func NewTodoRoute(e *echo.Echo, t TodoControllerInterface) *echo.Group {
	g := e.Group("/todos")
	g.GET("", t.GetTodos)
	g.GET("/:id", t.GetTodo)
	g.POST("/", t.CreateTodo)
	g.PUT("/:id", t.UpdateTodo)
	g.DELETE("/:id", t.DeleteTodo)

	return g
}

func NewWebTodoRoute(e *echo.Echo) *echo.Group {
	c := NewTodoWebController()
	g := NewTodoRoute(e, c)

	return g
}
