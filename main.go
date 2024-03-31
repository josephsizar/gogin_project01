package main

import (
	"goveri/groups"
	"goveri/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	v1 := router.Group("/about")
	{

		groups.About(v1.Group("/"))
	}

	auth := router.Group("/user")
	{
		auth.Use(middlewares.UserAuth)

		groups.User(auth.Group("/"))
	}

	router.Run(":3000")

}
