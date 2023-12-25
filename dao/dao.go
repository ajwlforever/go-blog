package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-blog/config"
	"strconv"
)

var db *sqlx.DB

// 初始化数据库
func initMysql() (err error) {
	fmt.Println(config.Conf.Db.Ports[0])
	dsn := config.Conf.Db.Username + ":" + config.Conf.Db.Password
	dsn += "@tcp(" + config.Conf.Db.Server + ":" + strconv.Itoa(config.Conf.Db.Ports[0]) + ")/" + config.Conf.Db.Dbname + "?charset=utf8mb4&parseTime=True"

	db, err = sqlx.Connect("mysql", dsn)
	fmt.Printf("dsn: %s\n", dsn)
	if err != nil {
		fmt.Printf("connect server failed, error: %v", err)
		return
	}
	db.SetMaxOpenConns(config.Conf.Db.ConnMax)
	db.SetMaxIdleConns(10)
	return
}

func init() {
	initMysql()
}
