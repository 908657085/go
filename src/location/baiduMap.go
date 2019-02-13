package location

import (
	"db"
	"fmt"
)

func InsertLocation(dbw db.DbWorker, radius float64, direction int, latitude float64, longitude float64) {
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
