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
	if MangoUrl == "" {
		fmt.Println("MangoUrl is empty")
		MONGODB := os.Getenv("MONGODB")
		DBUser := os.Getenv("MONGOD_DB_USER")
		DBPass := os.Getenv("MONGOD_DB_PASS")
		MangoUrl = "mongodb://" + DBUser + ":" + DBPass + MONGODB + "heroku_1vxk1j6t"
	}
	fmt.Println(MangoUrl)
	//Connect to MangoDB MONGODB_URI
	client, err := mongo.NewClient(options.Client().ApplyURI(MangoUrl + "?retryWrites=false"))
	if err != nil {
		fmt.Println(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(context.Background())
	col := client.Database("heroku_1vxk1j6t").Collection("photos")
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
