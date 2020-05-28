package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type todo struct {
	ID        int    `json:"id"`
	Schedule  string `json:"schedule"`
	Priority  string `json:"priority"`
	TimeLimit string `json:"time_limit"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:0111@/todo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	db.SingularTable(true)

	http.Handle("/", http.FileServer(http.Dir("views")))
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			display(w, r)
		case http.MethodPost:
			register(w, r)
		case http.MethodDelete:
			remove(w, r)
		}
	})
	log.Println("start http server :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func register(w http.ResponseWriter, r *http.Request) {
	var body todo

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	result := db.Create(&body)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), 500)
		return
	}
}
func display(w http.ResponseWriter, r *http.Request) {
	var todoList []todo
	result := db.Find(&todoList)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), 500)
		return
	}
	if err := json.NewEncoder(w).Encode(&todoList); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
func remove(w http.ResponseWriter, r *http.Request) {
	var body todo
	var err error
	body.ID, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := db.Delete(&body).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
