package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type todo struct {
	ID        int    `json:"id"`
	Schedule  string `json:"schedule"`
	Priority  string `json:"priority"`
	TimeLimit string `json:"timeLimit"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:0111@/todo")
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	result := db.Create(&body)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
func display(w http.ResponseWriter, r *http.Request) {
	var todoList []todo
	result := db.Find(&todoList)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	if err := json.NewEncoder(w).Encode(&todoList); err != nil {
		log.Fatal(err)
	}

}
func remove(w http.ResponseWriter, r *http.Request) {
	var body todo
	var err error
	body.ID, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	result := db.Delete(&body)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
