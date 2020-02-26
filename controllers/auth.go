package controllers

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/joho/godotenv"
)

func Secret(user, realm string) string {
	fmt.Println(fmt.Sprintf("../.env", os.Getenv("GO_ENV")))
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	USERName := os.Getenv("USER_NAME")
	USERPass := os.Getenv("USER_PASSWORD")

	fmt.Println(USERName)

	if USERName == "" {
		fmt.Println("Error loading USER_NAME")
	}
	if USERPass == "" {
		fmt.Println("Error loading USER_PASSWORD")
	}
	if user == USERName {
		// password is "hello"
		fmt.Println(USERPass)
		return USERPass
	}

	return ""
}

type MainController struct {
	beego.Controller
}

// func (this *MainController) Prepare() {
// 	a := auth.NewBasicAuthenticator("example.com", Secret)

// 	if username := a.CheckAuth(this.Ctx.Request); username == "" {
// 		a.RequireAuth(this.Ctx.ResponseWriter, this.Ctx.Request)
// 	}
// }

func (this *MainController) Get() {
	this.TplName = "owner.tpl"
}
