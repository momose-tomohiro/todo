package repository

import (
	"todo/api/domain/model"

	"github.com/labstack/echo"
)

type TodoRepository interface {
	GetTodoList() ([]model.Todo, error)
	RegisterTodo(c echo.Context) error
	RemoveTodo(c echo.Context) error
}
