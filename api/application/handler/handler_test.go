package handler

import (
	"log"
	"testing"
	"todo/api/domain/model"

	"github.com/labstack/echo"
)

type TestTodoService struct {
	err error
}

func (t *TestTodoService) GetTodoService() ([]model.Todo, error) {
	if t.err != nil {
		return nil, t.err
	}
	return []model.Todo{
		{ID: 1, Schedule: "test", Priority: "高", TimeLimit: "2020-06-12"},
		{ID: 2, Schedule: "test2", Priority: "低", TimeLimit: "2020-06-12"},
	}, nil
}

func (t *TestTodoService) RegisterTodoService(c echo.Context) error {
	return t.err
}

func (t *TestTodoService) RemoveTodoService(c echo.Context) error {
	return t.err
}
func TestGetTodoList(t *testing.T) {
	testTodo := &TestTodoService{}
	test, err := testTodo.GetTodoService()
	if err != nil {
		t.Fatalf("error: %d", err)
	}
	log.Println(test)
	t.Logf(test[0].Priority)
}

func TestRegister(t *testing.T) {
	testTodo := &TestTodoService{}
	e := echo.New()
	err := testTodo.RegisterTodoService()
}
