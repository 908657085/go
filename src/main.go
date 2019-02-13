package main

import (
	"db"
)

func main() {
	//hello world go!
	//fmt.Println("hello, world!")

	//web demo
	//web.Demo()

	//mysql demo
	db.Demo()
	//db.Init()
	defer db.Destory()

}
