package controllers

import (
    "fmt"
	"io/ioutil"
	"os"
	"github.com/astaxie/beego"
)

type UploadController struct {
    beego.Controller
}


func (this *UploadController) UploadFile() { 
	fmt.Println("File Upload Endpoint Hit")
    // this.Data["json"] = employees2
    // this.ServeJSON()
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
	//Create storage path
	fmt.Println("Create dir")
	err = os.MkdirAll("./storage/gallery", 0755)
	if err != nil {
		fmt.Println(err)
	}
	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("./storage/gallery", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
    this.Ctx.ResponseWriter.WriteHeader(200)
	fmt.Fprintf(this.Ctx.ResponseWriter, "Successfully Uploaded File\n")
    // this.Ctx.Redirect(200, "../")
}