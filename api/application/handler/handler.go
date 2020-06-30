package handler

import (
	"fmt"
	"net/http"
	"todo/api/domain/service"

	"github.com/labstack/echo"
)

type TodoHandler interface {
	GetTodoList(c echo.Context) error
	RegisterTodo(c echo.Context) error
	RemoveTodo(c echo.Context) error
}

type todoHandler struct {
	todo service.TodoServiceInterface
}

func NewTodoHandler(todo service.TodoServiceInterface) TodoHandler {
	return &todoHandler{todo: todo}
}

func (t *todoHandler) GetTodoList(c echo.Context) error {
	fmt.Println("確認")
	todoList, err := t.todo.GetTodoService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todoList)
}

func (t *todoHandler) RegisterTodo(c echo.Context) error {
	err := t.todo.RegisterTodoService(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}

func (t *todoHandler) RemoveTodo(c echo.Context) error {
	err := t.todo.RemoveTodoService(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}
