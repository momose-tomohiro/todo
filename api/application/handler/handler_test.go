package handler

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"todo/api/domain/model"

	"github.com/labstack/echo"
)

type TestTodoService struct {
	err error
}

func (t *TestTodoService) GetTodoService() ([]model.Todo, error) {
	t.err = errors.New("test error")
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
	var body model.Todo
	var err error

	body.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	return t.err
}
func TestGetTodoList(t *testing.T) {
	testTodo := &TestTodoService{}
	test, err := testTodo.GetTodoService()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(test)
	t.Logf("OK")
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
		t.Fatal(err)
	}
}

func TestRemoveTodo(t *testing.T) {
	testTodo := &TestTodoService{}
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := testTodo.RemoveTodoService(c)
	if err != nil {
		t.Fatal(err)
	}
}
