package web

import (
	"db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"upload"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!")
}

func Demo() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world! ")
	fmt.Fprintf(w, "go index page!")
}

//上传位置信息
func UploadLocation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var radius, latitude, longitude float64
	var direction int
	for k, v := range r.Form {
		switch k {
		case "radius":
			radius, _ = strconv.ParseFloat(strings.Join(v, ""), 32)
		case "direction":
			direction, _ = strconv.Atoi(strings.Join(v, ""))
		case "latitude":
			latitude, _ = strconv.ParseFloat(strings.Join(v, ""), 32)
		case "longitude":
			longitude, _ = strconv.ParseFloat(strings.Join(v, ""), 32)
		}
	}
	if latitude != 0 && longitude != 0 {
		id, err := db.InsertLocation(db.Dbw, radius, direction, latitude, longitude)
		if nil == err {
			io.WriteString(w, "upload location success"+strconv.Itoa(int(id)))
			return
		}
	}
	io.WriteString(w, "upload location error")
}

//查询最新位置信息
func QueryCurrentLocation(w http.ResponseWriter, r *http.Request) {
	locations, err := db.ListLocation(db.Dbw)
	fmt.Println(locations, err)
}

//查询路线信息
func ListLineMap(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var userName, password string
	apiReuslt := ApiResult{
		false,
		"",
		nil,
	}
	for k, v := range r.Form {
		switch k {
		case "userName":
			userName = strings.Join(v, "")
		case "password":
			password = strings.Join(v, "")
		}
	}
	if userName == "" || password == "" {
		io.WriteString(w, "login fail!")
		return
	}
	userInfo, err := db.Login(userName, password)
	if nil != err {
		fmt.Fprintf(w, "login fail! ", err)
		return
	}
	apiReuslt.Data = userInfo
	result, _ := json.Marshal(apiReuslt)
	fmt.Fprintf(w, string(result))
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var userName, password, nickName = r.FormValue("userName"), r.FormValue("password"), r.FormValue("nickName")
	tel, err := strconv.ParseInt(r.FormValue("tel"), 10, 64)
	if nil != err {
		fmt.Println("register user tel error: ", err)
		tel = 0
	}
	if userName == "" || password == "" {
		fmt.Fprintf(w, "register user fail! username or password nil!")
		fmt.Fprintf(w, " test123afa")
		return
	}
	userinfo, err := db.InsertUserInfo(userName, password, nickName, tel)
	if nil != err {
		fmt.Fprintf(w, "register user error", err)
		return
	}
	result, _ := json.Marshal(userinfo)
	fmt.Fprintf(w, string(result))
}

func Init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/uploadLocation", UploadLocation)
	http.HandleFunc("/queryCurrentLocation", QueryCurrentLocation)
	http.HandleFunc("/listLineMap", ListLineMap)
	http.HandleFunc("/registerUser", registerUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload.UploadHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
		return
	}
	fmt.Println("服务器启动成功")
}
