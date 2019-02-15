package LocationServer

import (
	"LocationServer/db"
	"LocationServer/web"
)

func main() {
	//hello world go!
	//fmt.Println("hello, world!")

	//web demo
	//web.Demo()

	//mysql demo
	db.Init()
	defer db.Destory()
	web.Init()
}
