package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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
	body := new(model.Todo)
	if err := c.Bind(body); err != nil {
		return err
	}
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

func TestRegisterTodo(t *testing.T) {
	testTodo := &TestTodoService{}
	testJSON := `{"schedule":"test3","priority":"中","time_limit":"2020-6-12"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(testJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := testTodo.RegisterTodoService(c)
	if err != nil {
		t.Fatalf("error: %d", err)
	}
}

func TestRemoveTodo(t *testing.T) {
	testTodo := &TestTodoService{}
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := testTodo.RemoveTodoService(c)
	if err != nil {
		t.Fatalf("error: %d", err)
	}
}
