package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	upload_path string = "./upload_file/"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//上传
func UploadHandle(w http.ResponseWriter, r *http.Request) {
	//从请求当中判断方法
	if r.Method == "GET" {
		io.WriteString(w, "<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传图片</label>"+":"+
			"<input type=\"file\" name='file'  /><br/><br/>    "+
			"<label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		r.ParseMultipartForm(32 << 20)
		//获取上传token
		params := r.FormValue("params")
		fmt.Println(params)
		isEncrypted := r.FormValue("isEncrypted")
		fmt.Println(isEncrypted)
		timeStamp := r.FormValue("timeStamp")
		fmt.Println(timeStamp)
		randomNum := r.FormValue("randomNum")
		fmt.Println(randomNum)
		sign := r.FormValue("sign")
		fmt.Println(sign)
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//创建文件
		fmt.Println(head.Filename)
		exists, err := PathExists(upload_path)
		if err != nil {
			fmt.Println("检测文件目录失败")
			return
		}
		if !exists {
			fmt.Println("upload_path not exists,mkdir")
			os.Mkdir(upload_path, os.ModePerm)
		}
		fW, err := os.Create(upload_path + head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		fmt.Println("文件上传成功!")
		io.WriteString(w, head.Filename+"params: "+params+"isEncrypted: "+isEncrypted+" timeStamp: "+timeStamp+" randomNum: "+randomNum+" sign : "+sign)
	}
}
