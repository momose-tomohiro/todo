package service

import (
	"todo/api/domain/model"
	"todo/api/domain/repository"

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

func (t *todoServiceStruct) GetTodoService() ([]model.Todo, error) {
	return t.todoRepo.GetTodoList()
}

func (t *todoServiceStruct) RegisterTodoService(c echo.Context) error {
	return t.todoRepo.RegisterTodo(c)
}

func (t *todoServiceStruct) RemoveTodoService(c echo.Context) error {
	return t.todoRepo.RemoveTodo(c)
}
