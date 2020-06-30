package infrastructure

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:tomoaki7@tcp(172.21.0.3:3306)/todo")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	Engine.SetMapper(core.GonicMapper{})
}
