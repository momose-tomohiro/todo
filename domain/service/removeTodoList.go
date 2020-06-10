package service

import "github.com/labstack/echo"

func (t *todoServiceStruct) RemoveTodoService(c echo.Context) error {
	return t.todoRepo.RemoveTodo(c)
}
