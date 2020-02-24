package routers

import (
	"firstBee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.FirstController{},"get:IndexPage")
	beego.Router("/employees", &controllers.FirstController{}, "post:GetEmployees")
    beego.Router("/employees", &controllers.FirstController{}, "get:GetEmployees")
	// beego.Router("/dashboard", &controllers.FirstController{}, "get:Dashboard")
}
