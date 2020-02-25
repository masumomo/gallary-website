package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
    "io/ioutil"
    "bufio"
    "encoding/base64"
    "os"
)

type FirstController struct {
    beego.Controller
}

type Photo struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Src string `json:"src"`
}

type Photos []Photo

var photos []Photo

func init() {
    files, _ := ioutil.ReadDir("./storage/gallery")
    files1, _ := ioutil.ReadDir(".")
    files2, _ := ioutil.ReadDir("./storage")
    fmt.Println(files)
    fmt.Println(files1)
    fmt.Println(files2)
    for _, file := range files {
        filename := file.Name()
        f, err := os.Open("./storage/gallery/" + filename) 
        reader := bufio.NewReader(f)
        content, _ := ioutil.ReadAll(reader)
        encoded := base64.StdEncoding.EncodeToString(content)
        defer f.Close()
        if err != nil {
            return
        }
        fmt.Println(filename)
        photos = append(photos,Photo{ID: 1, Name: filename, Src: encoded})

    }
}

func (this *FirstController) GetPhotos() {
    this.Ctx.ResponseWriter.WriteHeader(200)
    this.Data["json"] = photos
    this.ServeJSON()
}

func (this *FirstController) IndexPage() {
    this.Data["Photos"] = photos
    this.TplName = "index.tpl"
}
