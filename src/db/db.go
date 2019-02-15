package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

type DbWorker struct {
	Dsn string
	Db  *sql.DB
}

var Dbw DbWorker

func Init() {
	var err error
	var timeLocation = url.QueryEscape("Asia/Shanghai")
	Dbw = DbWorker{
		Dsn: "root:devil@tcp(101.132.38.194:3306)/baiduMap?charset=utf8mb4&loc" + timeLocation + "&parseTime=true",
	}
	Dbw.Db, err = sql.Open("mysql", Dbw.Dsn)
	if err != nil {
		fmt.Println("database open error: ", err)
		panic(err)
		return
	}
	fmt.Println("database open: ", Dbw.Db)
}

func Destory() {
	fmt.Println("database close")
	Dbw.Db.Close()
}

func (dbw *DbWorker) Check() error {
	if nil == Dbw.Db {
		return errors.New("db connection nil")
	}
	return nil
}
