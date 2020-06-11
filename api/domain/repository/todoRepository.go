package repository

import (
	"todo/domain/model"

	"github.com/labstack/echo"
)

type TodoRepository interface {
	GetTodoListAll() ([]model.Todo, error)
	RegisterTodo(c echo.Context) error
	RemoveTodo(c echo.Context) error
}
