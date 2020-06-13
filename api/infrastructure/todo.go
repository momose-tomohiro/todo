package infrastructure

import (
	"strconv"
	"todo/api/domain/model"
	"todo/api/domain/repository"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type todoInfraStruct struct {
	db *gorm.DB
}

func NewTodoDB(db *gorm.DB) repository.TodoRepository {
	return &todoInfraStruct{db: db}
}

func (t *todoInfraStruct) GetTodoList() ([]model.Todo, error) {
	var todoList []model.Todo
	result := t.db.Find(&todoList)
	return todoList, result.Error
}

func (t *todoInfraStruct) RegisterTodo(c echo.Context) error {
	body := new(model.Todo)
	if err := c.Bind(body); err != nil {
		return err
	}

	result := t.db.Create(&body)
	return result.Error

}

func (t *todoInfraStruct) RemoveTodo(c echo.Context) error {
	var body model.Todo
	var err error

	body.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = t.db.Delete(&body).Error
	return err
}
