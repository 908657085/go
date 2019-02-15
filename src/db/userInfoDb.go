package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserInfo struct {
	Id            int
	UserName      sql.NullString
	Password      sql.NullString
	NickName      sql.NullString
	Tel           sql.NullInt64
	CreatTime     time.Time
	LastLoginTime time.Time
}

func InsertUserInfo(userName, password, nickName string, tel int64) (user UserInfo, err error) {
	var userInfo UserInfo
	if e := Dbw.Check(); e != nil {
		return userInfo, e
	}
	stmt, err := Dbw.Db.Prepare(`INSERT INTO user_info (user_name, password,nick_name,tel) VALUES (?,?,?,?)`)
	defer stmt.Close()
	if err != nil {
		fmt.Println("insert userInfo prepare error :", err)
		return userInfo, err
	}
	ret, err := stmt.Exec(userName, password, nickName, tel)
	if err != nil {
		fmt.Printf("insert userInfo error: %v\n", err)
		return userInfo, err
	}
	LastInsertId, Err := ret.LastInsertId()
	if nil != Err {
		fmt.Println("get userInfo LastInsertId error:", err)
		return userInfo, err
	}
	userInfo = UserInfo{
		int(LastInsertId),
		sql.NullString{userName, true},
		sql.NullString{"", true},
		sql.NullString{nickName, true},
		sql.NullInt64{tel, true},
		time.Now().UTC(),
		time.Now().UTC(),
	}
	return userInfo, nil
}

func UpdateUserInfo(id int64, nickName string, tel int64) (updateResult bool, err error) {
	if e := Dbw.Check(); e != nil {
		return false, e
	}
	stmt, err := Dbw.Db.Prepare(` UPDATE user_info SET nick_name=?,tel=? WHERE id = ?`)
	defer stmt.Close()
	if nil != err {
		fmt.Printf("update userInfo prepare error: %v\n", err)
	}
	ret, err := stmt.Exec(nickName, tel, id)
	if nil != err {
		fmt.Printf("update userInfo error: %v\n", err)
	}
	rowsAffected, err := ret.RowsAffected()
	if nil != err {
		fmt.Printf("update userInfo rowsAffected error: %v\n", err)
	}
	if rowsAffected > 0 {
		return true, nil
	}
	return false, err
}

func Login(userName, password string) (user UserInfo, err error) {
	var userInfo UserInfo
	rows, err := Dbw.Db.Query(`SELECT id,user_name,nick_name,tel,last_login_time FROM user_info where (user_name =? AND password = ?)`, userName, password)
	defer rows.Close()

	if nil != err {
		fmt.Printf("query userInfo error: %v\n", err)
		return userInfo, err
	}
	columns, err := rows.Columns()
	if len(columns) > 0 {
		for rows.Next() {
			rows.Scan(&userInfo.Id, &userInfo.UserName, &userInfo.NickName, &userInfo.Tel, &userInfo.LastLoginTime)
		}
		return userInfo, nil
	}
	return userInfo, errors.New("username or password error!")
}

func QueryUserInfo() UserInfo {
	return UserInfo{}
}
