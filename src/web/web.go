package web

import (
	"db"
	"fmt"
	"io"
	"location"
	"log"
	"net/http"
	"strconv"
	"strings"
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

//上传位置信息
func UploadLocation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var radius, latitude, longitude float64
	var direction int
	for k, v := range r.Form {
		switch k {
		case "radius":
			radius, _ = strconv.ParseFloat(strings.Join(v, ""), 32)
			break
		case "direction":
			direction, _ = strconv.Atoi(strings.Join(v, ""))
			break
		case "latitude":
			latitude, _ = strconv.ParseFloat(strings.Join(v, ""), 32)
			break
		case "longitude":
			longitude, _ = strconv.ParseFloat(strings.Join(v, ""), 32)
			break
		}
	}
	if latitude != 0 && longitude != 0 {
		id, err := location.InsertLocation(db.Dbw, radius, direction, latitude, longitude)
		if nil == err {
			io.WriteString(w, "upload location success"+strconv.Itoa(int(id)))
			return
		}
	}
	io.WriteString(w, "upload location error")
}

//查询最新位置信息
func QueryCurrentLocation(w http.ResponseWriter, r *http.Request) {
	locations, err := location.ListLocation(db.Dbw)
	fmt.Println(locations, err)
}

//查询路线信息
func ListLineMap(w http.ResponseWriter, r *http.Request) {

}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world! ")
	fmt.Fprintf(w, "go index page!")
}

func Init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/uploadLocation", UploadLocation)
	http.HandleFunc("/queryCurrentLocation", QueryCurrentLocation)
	http.HandleFunc("/listLineMap", ListLineMap)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
		return
	}
	fmt.Println("服务器启动成功")
}
