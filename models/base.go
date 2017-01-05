package models

import (
	// "database/sql"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:digitalx168@tcp(127.0.0.1:3306)/task?charset=utf8mb4")
}
