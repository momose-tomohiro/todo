package infrastructure

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:tomoaki7@/todo")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	DB.SingularTable(true)
}
