package infrastructure

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"xorm.io/core"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:tomoaki7@/todo")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	Engine.SetMapper(core.GonicMapper{})
}
