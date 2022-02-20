package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"GoGetEmployed/controllers"
	// "GoGetEmployed/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/static", "./public")

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	router.LoadHTMLGlob("./templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "GoGetEmployed",
			"indexActive": true,
		})
	})

	router.GET("/jobs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "job-leads.tmpl", gin.H{
			"title": "Main website",
			"jobLeadsActive": true,
		})
	})

	// router.Use(middlewares.AuthMiddleware())
	//
	// v1 := router.Group("v1")
	// {
	// 	userGroup := v1.Group("user")
	// 	{
	// 		user := new(controllers.UserController)
	// 		userGroup.GET("/:id", user.Retrieve)
	// 	}
	// }
	return router

}
