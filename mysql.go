package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"time"
)

type DbWorker struct {
	Dsn      string
	Db       *sql.DB
	Location location
}
type location struct {
	Id         int
	Radius     sql.NullFloat64
	Direction  sql.NullInt64
	Latitude   sql.NullFloat64
	Longitude  sql.NullFloat64
	createTime time.Time
}

func main() {
	var err error
	var timeLocation = url.QueryEscape("Asia/Shanghai")
	dbw := DbWorker{
		Dsn: "root:devil@tcp(101.132.38.194:3306)/baiduMap?charset=utf8mb4&loc" + timeLocation + "&parseTime=true",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(dbw.Db)
	defer dbw.Db.Close()

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

func (dbw *DbWorker) QueryDataPre() {
	dbw.Location = location{}
}
func (dbw *DbWorker) queryData() {
	stmt, _ := dbw.Db.Prepare(`SELECT * From location `)
	defer stmt.Close()

	dbw.QueryDataPre()

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&dbw.Location.Id, &dbw.Location.Radius, &dbw.Location.Direction, &dbw.Location.Latitude, &dbw.Location.Longitude, &dbw.Location.createTime)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		fmt.Println("get data, id: ", dbw.Location.Id,
			" radius: ", float64(dbw.Location.Radius.Float64),
			" direction: ", int(dbw.Location.Direction.Int64),
			" latitude: ", float64(dbw.Location.Latitude.Float64),
			" longitude: ", float64(dbw.Location.Longitude.Float64),
			" createTime: ", dbw.Location.createTime.String())
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
