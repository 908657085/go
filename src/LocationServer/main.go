package main

import (
	"LocationServer/db"
	"LocationServer/web"
)

func main() {
	//hello world go!cd
	//fmt.Println("hello, world!")

	//web demo
	//web.Demo()

	//mysql demo
	db.Init()
	defer db.Destroy()
	web.Init()
}
