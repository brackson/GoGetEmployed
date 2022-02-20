package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"GoGetEmployed/config"
	"gorm.io/gorm"
	// _ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	// "github.com/k0kubun/pp"

	// "GoGetEmployed/controllers"
	// "GoGetEmployed/models"
	// "GoGetEmployed/db"
)

type Job struct {
	JobId			int64 `db:"job_id" form:"JobId" gorm:"primaryKey;auto_increment;not_null"`
	CompanyName		string `db:"company_name" form:"CompanyName"`
	Url				string `db:"url" form:"Url"`
	Description		string `db:"description" form:"Description"`
	Notes			string `db:"notes" form:"Notes"`
	StatusId		int64 `db:"status_id" form:"StatusId"`
	Role			string `db:"role" form:"Role"`
}

type Task struct {
	TaskId			string
	Description	string
	State				int64
	JobId		int64
}

// type JobWithTasks struct {
// 	JobId			string
// 	CompanyName	string
// 	Url				string
// 	Description		string
// 	Notes			string
// 	StatusId		int64
// 	Role			string
// 	Tasks			[]Task
// }

var db *gorm.DB
var err error

func Init() {
	config := config.GetConfig()

	db, err := gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Job{})
	db.AutoMigrate(&Task{})

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/static", "./public")

	router.LoadHTMLGlob("./templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "GoGetEmployed",
			"indexActive": true,
		})
	})

	router.GET("/jobs", func(c *gin.Context) {
		var jobs []Job
		db.Find(&jobs)

		c.HTML(http.StatusOK, "job-leads.tmpl", gin.H{
			"title": "Main website",
			"jobLeadsActive": true,
			"jobs": jobs,
		})
	})

	router.GET("/jobs/update", func(c *gin.Context) {
		var job Job
		id := c.Query("JobId")

		if err := db.First(&job, id).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.BindQuery(&job)

		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&job)
		c.Redirect(302, "/jobs")
	})

	router.GET("/jobs/new", func(c *gin.Context) {
		var job Job
		c.BindQuery(&job)

		db.Debug().Create(&job)
		c.Redirect(302, "/jobs")
	})

	router.Run(config.GetString("server.port"))
}
