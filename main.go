package main

import (
	_ "firstBee/routers"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}
	beego.Run()
}
