package routers

import (
	"firstBee/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.FirstController{}, "get:IndexPage")
	beego.Router("/login", &controllers.MainController{})
	beego.Router("/upload", &controllers.UploadController{}, "post:UploadFile")
	beego.Router("/photos", &controllers.FirstController{}, "post:GetPhotos")
}
