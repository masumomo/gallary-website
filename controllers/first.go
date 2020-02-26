package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FirstController struct {
	beego.Controller
}

type Photo struct {
	Name string
	Src  string
	Date time.Time
}

type Photos []Photo

var photos []Photo

func init() {
	// cwd, err := os.Getwd()
	// fmt.Println(cwd)
	// fmt.Println(err)
	fmt.Println(fmt.Sprintf("../.env", os.Getenv("GO_ENV")))
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	MangoUrl := os.Getenv("MONGODB_URI")
	MONGODB := os.Getenv("MONGODB")
	DBUser := os.Getenv("MONGODB_USER")
	DBPass := os.Getenv("MONGODB_PASS")
	DBName := os.Getenv("MONGODB_NAME")
	if MangoUrl == "" {
		fmt.Println("MangoUrl is empty")
		MangoUrl = "mongodb://" + DBUser + ":" + DBPass + MONGODB + DBName
	}
	fmt.Println(MangoUrl)
	//Connect to MangoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(MangoUrl + "?retryWrites=false"))
	if err != nil {
		fmt.Println(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(context.Background())
	col := client.Database(DBName).Collection("photos")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := col.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Photo
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}

		photos = append(photos, result)
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
