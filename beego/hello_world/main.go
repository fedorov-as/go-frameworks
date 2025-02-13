package main

import (
	_ "hello_world/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

