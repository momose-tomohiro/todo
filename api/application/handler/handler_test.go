package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"todo/api/domain/model"

	"github.com/labstack/echo"
)

type todoService struct {
	err error
}

func (t *todoService) GetTodoService() ([]model.Todo, error) {
	//t.err = errors.New("test error")
	if t.err != nil {
		return nil, t.err
	}
	return []model.Todo{
		{ID: 1, Schedule: "test", Priority: "高", TimeLimit: "2020-06-12"},
	}, nil
}

func (t *todoService) RegisterTodoService(c echo.Context) error {
	body := new(model.Todo)
	if err := c.Bind(body); err != nil {
		return err
	}
	return t.err
}

func (t *todoService) RemoveTodoService(c echo.Context) error {
	var body model.Todo
	var err error

	body.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	return t.err
}

func TestGetTodoList(t *testing.T) {
	serviceErr := errors.New("test error")
	tests := []struct {
		serviceErr error
		endpoint   string
		wantStatus int
		wantBody   string
	}{
		{nil, "/todos", 200, `[{"id":1,"schedule":"test","priority":"高","time_limit":"2020-06-12"}]`},
		{serviceErr, "/todos", 500, `{"message":"cannot get users: service layer error"}`},
	}

	var h TodoHandler
	for i, tt := range tests {
		if tt.serviceErr != nil {
			h = NewTodoHandler(&todoService{serviceErr})
		} else {
			h = NewTodoHandler(&todoService{})
		}

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, tt.endpoint, nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if err := h.GetTodoList(c); err != nil {
			t.Fatalf("case:[%d] GET %s error:%s", i, tt.endpoint, err)
		}

		if rec.Code != tt.wantStatus {
			t.Fatalf("case:[%d] GET %s unexpected status code, want=%d, got=%d", i, tt.endpoint, tt.wantStatus, rec.Code)
		}
		gotBody := strings.Trim(rec.Body.String(), "\n\r")
		if gotBody != tt.wantBody {
			t.Fatalf("case:[%d] GET %s unexpected body, want=%s, got=%s", i, tt.endpoint, tt.wantBody, gotBody)
		}
	}
	t.Logf("OK")
}

/*
func TestGetTodoList(t *testing.T) {
	serviceErr := errors.New("test error")
	tests := []struct {
		serviceErr error
		endpoint   string
		wantStatus int
		wantBody   string
	}{
		{nil, "/todos", 200, `[{"id":1,"Schedule":"test","Priority":"高","TimeLimit","2020-06-16"}]`},
	}

	testTodo := &TodoService{}
	test, err := testTodo.GetTodoService()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(test)
	t.Logf("OK")
}

func TestRegisterTodo(t *testing.T) {
	testTodo := &TodoService{}
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
	t.Logf("OK")
}

func TestRemoveTodo(t *testing.T) {
	testTodo := &TodoService{}
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
	t.Logf("OK")
}
*/
