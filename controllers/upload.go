package controllers

import (
    "fmt"
	"io/ioutil"
	"github.com/astaxie/beego"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
    "bufio"
	"encoding/base64"
)

type UploadController struct {
    beego.Controller
}





func (this *UploadController) UploadFile() { 

	fmt.Println("File Upload Endpoint Hit")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	this.Ctx.Request.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err :=  this.Ctx.Request.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fileTitle :=  handler.Filename;
	// tempFile, err := ioutil.TempFile("", "upload-*.png")

	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)

	defer file.Close()
	if err != nil {
		return
	}

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//MONGODB_URI
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mitsuru:masumotomitsuru0108@ds133137.mlab.com:33137/heroku_1vxk1j6t?retryWrites=false"))
    if err != nil {
		fmt.Println(err)
    }
    if err = client.Connect(context.Background()); err != nil {
		fmt.Println(err)
    }
    defer client.Disconnect(context.Background())

	col := client.Database("heroku_1vxk1j6t").Collection("photos")
	
	doc := Photo{
        fileTitle,
        encoded,
        time.Now(),
    }
	if  _, err =  col.InsertOne(context.Background(), doc); err != nil {
		fmt.Println(err)
    }
    this.Ctx.ResponseWriter.WriteHeader(200)
	fmt.Fprintf(this.Ctx.ResponseWriter, "Successfully Uploaded File\n")
}