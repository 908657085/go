package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Location struct {
	Id         int
	UId        int
	Radius     sql.NullFloat64
	Direction  sql.NullInt64
	Latitude   sql.NullFloat64
	Longitude  sql.NullFloat64
	createTime time.Time
}

func InsertLocation(dbw DbWorker, radius float64, direction int, latitude float64, longitude float64) (lastInsertId int64, err error) {
	stmt, err := dbw.Db.Prepare(`INSERT INTO location (radius, direction,latitude,longitude) VALUES (?,?,?,?)`)
	defer stmt.Close()
	if err != nil {
		fmt.Println("insert location prepare error :", err)
		return
	}
	ret, err := stmt.Exec(radius, direction, latitude, longitude)
	if err != nil {
		fmt.Printf("insert location error: %v\n", err)
		return
	}
	LastInsertId, Err := ret.LastInsertId()
	if nil != Err {
		fmt.Println("get location LastInsertId error:", err)
		return 0, err
	}
	//if rowsAffected, err = ret.RowsAffected(); nil == err {
	//	fmt.Println("RowsAffected:", rowsAffected)
	//
	//}
	return LastInsertId, nil
}

func ListLocation(dbw DbWorker) (locations []Location, err error) {
	stmt, _ := dbw.Db.Prepare(`SELECT id,radius,direction,latitude,longitude From location ORDER BY create_time DESC limit 0,1`)
	defer stmt.Close()

	var location Location

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Printf("list location error: %v\n", err)
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
	return nil, nil
}
