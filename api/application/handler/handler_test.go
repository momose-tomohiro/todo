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
}

func TestRegisterTodo(t *testing.T) {
	serviceErr := errors.New("test error")
	tests := []struct {
		serviceErr error
		endpoint   string
		wantStatus int
		reqBody    string
	}{
		{nil, "/todos", 200, `{"schedule":"test","priority":"高","time_limit":"2020-06-16"}`},
		{nil, "/todos", 200, `{"schedule":"test2","priority":"低","time_limit":"未定"}`},
	}
	var h TodoHandler
	for i, tt := range tests {
		if tt.serviceErr != nil {
			h = NewTodoHandler(&todoService{serviceErr})
		} else {
			h = NewTodoHandler(&todoService{})
		}

		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, tt.endpoint, strings.NewReader(tt.reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if err := h.RegisterTodo(c); err != nil {
			t.Fatalf("case:[%d] POST %s error:%s", i, tt.endpoint, err)
		}

		if rec.Code != tt.wantStatus {
			t.Fatalf("case:[%d] POST %s unexpected status code, want=%d, got=%d", i, tt.endpoint, tt.wantStatus, rec.Code)
		}
	}
}

func TestRemoveTodo(t *testing.T){
	serviceErr := errors.New("test error")
	tests := []struct {
		serviceErr error
		endpoint   string
		wantStatus int
		wantBody    string
	}{
		{nil, "/todos/:id", 200, `"OK"`},
		{nil, "/todos/:id", 200, `"OK"`},
		{serviceErr, "/todos/:id", 500, `{"message":"cannot get id:3 : service layer error"}`},
		
	}
	var h TodoHandler
	for i, tt := range tests {
		if tt.serviceErr != nil {
			h = NewTodoHandler(&todoService{serviceErr})
		} else {
			h = NewTodoHandler(&todoService{})
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(i + 1))
		if err := h.RemoveTodo(c); err != nil {
			t.Fatalf("case:[%d] POST %s error:%s", i, tt.endpoint, err)
		}
		if rec.Code != tt.wantStatus {
			t.Fatalf("case:[%d] POST %s unexpected status code, want=%d, got=%d", i, tt.endpoint, tt.wantStatus, rec.Code)
		}
		gotBody := strings.Trim(rec.Body.String(), "\n\r")
		if gotBody != tt.wantBody {
			t.Fatalf("case:[%d] GET %s unexpected body, want=%s, got=%s", i, tt.endpoint, tt.wantBody, gotBody)
		}
		
	}
}

/*

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
