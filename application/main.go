package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"todo/application/handler"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type todo struct {
	ID        int    `json:"id"`
	Schedule  string `json:"schedule"`
	Priority  string `json:"priority"`
	TimeLimit string `json:"time_limit"`
}

var db *gorm.DB

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	test := handler.Display()
	log.Println(test)

	var err error
	db, err = gorm.Open("mysql", "root:tomoaki7@/todo")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer db.Close()
	db.SingularTable(true)

	todos := e.Group("/todos")
	todos.GET("", display)
	todos.POST("", register)
	todos.DELETE("/:id", remove)

	e.Logger.Fatal(e.Start(":8888"))
}
func display(c echo.Context) error {
	var todoList []todo

	result := db.Find(&todoList)
	if result.Error != nil {
		//500番
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, todoList)
}

func register(c echo.Context) error {
	body := new(todo)

	if err := c.Bind(body); err != nil {
		//400番
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := db.Create(&body)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, "OK")
}

func remove(c echo.Context) error {
	var body todo
	var err error

	body.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := db.Delete(&body).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, "OK")
}
