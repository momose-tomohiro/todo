package service

import (
	"todo/domain/model"
	"todo/domain/repository"

	"github.com/labstack/echo"
)

type todoServiceStruct struct {
	todoRepo repository.TodoRepository
}

type TodoServiceInterface interface {
	GetTodoService() ([]model.Todo, error)
	RegisterTodoService(c echo.Context) error
	RemoveTodoService(c echo.Context) error
}

func NewTodoService(t repository.TodoRepository) TodoServiceInterface {
	return &todoServiceStruct{todoRepo: t}
}