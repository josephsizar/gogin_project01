package controllers

import "github.com/gin-gonic/gin"

func AboutIndex(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome Again in about index",
	})
}
