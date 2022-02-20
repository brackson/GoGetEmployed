package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"GoGetEmployed/config"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/k0kubun/pp"

	// "GoGetEmployed/controllers"
	// "GoGetEmployed/models"
	// "GoGetEmployed/db"
)

type Job struct {
	JobId			string
	CompanyName	string
	Url				string
	Description		string
	Notes			string
	StatusId		int64
}

type Task struct {
	TaskId			string
	Description	string
	State				int64
	JobId		int64
}

var db *gorm.DB
var err error

func Init() {
	config := config.GetConfig()
	// r := NewRouter()

	db, err = gorm.Open("sqlite3", "./data.db")
	//db, _ = gorm.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Job{})
	db.AutoMigrate(&Task{})

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/static", "./public")

	// health := new(controllers.HealthController)
	// router.GET("/health", health.Status)

	// db := db.GetDB()

	router.LoadHTMLGlob("./templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "GoGetEmployed",
			"indexActive": true,
		})
	})

	var jobs []Job
	db.Find(&jobs)
	pp.Print(jobs)

	router.GET("/jobs", func(c *gin.Context) {

		c.HTML(http.StatusOK, "job-leads.tmpl", gin.H{
			"title": "Main website",
			"jobLeadsActive": true,
			"jobs": jobs,
		})
	})

	router.Run(config.GetString("server.port"))
}
