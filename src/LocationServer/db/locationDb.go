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
	Addr       sql.NullString
	createTime time.Time
}

type UserLocation struct {
	UserId    int
	UserName  string
	NickName  string
	Tel       string
	Radius    float64
	Direction int
	Latitude  float64
	Longitude float64
	Addr      string
}

func InsertLocation(userId int, radius float64, direction int, latitude float64, longitude float64, addr string) (lastInsertId int64, err error) {
	if e := Dbw.Check(); e != nil {
		return 0, e
	}
	stmt, err := Dbw.Db.Prepare(`INSERT INTO location (u_id,radius, direction,latitude,longitude,addr) VALUES (?,?,?,?,?,?)`)
	defer stmt.Close()
	if err != nil {
		fmt.Println("insert location prepare error :", err)
		return
	}
	ret, err := stmt.Exec(userId, radius, direction, latitude, longitude, addr)
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

func QueryCurrentLocation() (userLocations []UserLocation, err error) {
	if e := Dbw.Check(); e != nil {
		return nil, e
	}
	stmt, _ := Dbw.Db.Prepare(`select c.id,c.user_name,c.nick_name,c.tel,b.radius,b.direction,b.latitude,b.longitude,b.addr from (select * from (select * from location order by create_time desc) a where (u_id is not null and create_time is not null) group by u_id order by u_id asc) b ,user_info c where b.u_id=c.id`)
	defer stmt.Close()

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Printf("query current location error: %v\n", err)
		return nil, err
	}

	var userLocation UserLocation
	var data []UserLocation

	for rows.Next() {
		err := rows.Scan(&userLocation.UserId, &userLocation.UserName, &userLocation.NickName, &userLocation.Tel, &userLocation.Radius, &userLocation.Direction, &userLocation.Latitude, &userLocation.Longitude, &userLocation.Addr)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		data = append(data, userLocation)
	}
	return data, nil
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
