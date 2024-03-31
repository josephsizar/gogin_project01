package controllers

import (
	"fmt"
	"goveri/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {

	age := c.Query("age")
	useragent := c.GetHeader("User-Agent")
	ipaddr := c.ClientIP()

	fmt.Println("🔥 Client Api " + utilities.Green + ipaddr + utilities.Reset)

	c.JSON(200, gin.H{
		"message":           "Root user Authorized " + age,
		"User-Agent":        "" + useragent,
		"Client IP Address": ipaddr,
	})
}

func UserAll(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "I am using HTML Templates",
	})
}

func UserById(c *gin.Context) {

	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "This User With Id : " + id,
	})
}
