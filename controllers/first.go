package controllers

import "github.com/astaxie/beego"

type FirstController struct {
    beego.Controller
}

type Employee struct {
    ID        int    `json:"id"`
    FirstName string `json:"firstName"`
    LastName  string `json:lastName`
}

type Employees []Employee

var employees []Employee

func init() {
    employees = Employees{
        Employee{ID: 1, FirstName: "Foo", LastName: "Bar"},
        Employee{ID: 2, FirstName: "Baz", LastName: "Qux"},
    }
}

func (this *FirstController) GetEmployees() {
    this.Ctx.ResponseWriter.WriteHeader(200)
    this.Data["json"] = employees
    this.ServeJSON()
}

func (this *FirstController) IndexPage() {
    this.Data["employees"] = employees
    this.TplName = "index.tpl"
}