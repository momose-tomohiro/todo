package service

import (
	"github.com/labstack/echo"
)

func (t *todoServiceStruct) RegisterTodoService(c echo.Context) error {
	return t.todoRepo.RegisterTodo(c)
}