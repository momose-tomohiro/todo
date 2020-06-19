package infrastructure

import (
	"strconv"
	"todo/api/domain/model"
	"todo/api/domain/repository"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

type todoInfraStruct struct {
	engine *xorm.Engine
}

func NewTodoDB(engine *xorm.Engine) repository.TodoRepository {
	return &todoInfraStruct{engine: engine}
}

func (t *todoInfraStruct) GetTodoList() ([]model.Todo, error) {
	var todoList []model.Todo
	err := t.engine.Find(&todoList)
	return todoList, err
}

func (t *todoInfraStruct) RegisterTodo(c echo.Context) error {
	body := new(model.Todo)
	if err := c.Bind(body); err != nil {
		return err
	}

	_, err := t.engine.Insert(body)
	return err

}

func (t *todoInfraStruct) RemoveTodo(c echo.Context) error {
	body := new(model.Todo)
	var err error

	body.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	_, err = t.engine.Where("id = ?", body.ID).Delete(body)
	return err
}
