package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"time"
)

type DbWorker struct {
	Dsn string
	Db  *sql.DB
}

type Location struct {
	Id         int
	Radius     sql.NullFloat64
	Direction  sql.NullInt64
	Latitude   sql.NullFloat64
	Longitude  sql.NullFloat64
	createTime time.Time
}

var dbw DbWorker

func Init() {
	var err error
	var timeLocation = url.QueryEscape("Asia/Shanghai")
	dbw = DbWorker{
		Dsn: "root:devil@tcp(101.132.38.194:3306)/baiduMap?charset=utf8mb4&loc" + timeLocation + "&parseTime=true",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(dbw.Db)
}

func Destory() {
	fmt.Println("database close")
	dbw.Db.Close()
}

func Demo() {
	Init()
	dbw.insertData(2, 2, 2, 2)
	dbw.queryData()
}

func (dbw *DbWorker) insertData(radius float64, direction int, latitude float64, longitude float64) {
	stmt, err := dbw.Db.Prepare(`INSERT INTO location (radius, direction,latitude,longitude) VALUES (?,?,?,?)`)
	defer stmt.Close()
	if err != nil {
		fmt.Println("insert prepare error :", err)
		return
	}
	ret, err := stmt.Exec(radius, direction, latitude, longitude)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

func (dbw *DbWorker) queryData() {
	stmt, _ := dbw.Db.Prepare(`SELECT * From location `)
	defer stmt.Close()

	var location Location

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&location.Id, &location.Radius, &location.Direction, &location.Latitude, &location.Longitude, &location.createTime)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		fmt.Println("get data, id: ", location.Id,
			" radius: ", float64(location.Radius.Float64),
			" direction: ", int(location.Direction.Int64),
			" latitude: ", float64(location.Latitude.Float64),
			" longitude: ", float64(location.Longitude.Float64),
			" createTime: ", location.createTime.String())
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
