package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"myblog_go/services/setting"
	"strings"
)

//Db数据库连接池
var Db *sql.DB

type DbConfig struct {
	host string
	port string
	user string
	pass string
	db   string
}

func init() {
	dcfg := DbConfig{
		setting.GetConfig("MYSQL", "host"),
		setting.GetConfig("MYSQL", "port"),
		setting.GetConfig("MYSQL", "user"),
		setting.GetConfig("MYSQL", "pass"),
		setting.GetConfig("MYSQL", "db"),
	}

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{dcfg.user, ":", dcfg.pass, "@tcp(", dcfg.host, ":", dcfg.port, ")/", dcfg.db, "?charset=utf8&parseTime=true"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	Db, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	Db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	Db.SetMaxIdleConns(10)
	//验证连接
	if err := Db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
	defer End()
}

func End() {
	//e := Db.Close()
	//if e != nil {
	//
	//}
}

func Query() {

}
