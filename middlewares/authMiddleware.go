package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	user := c.Query("user")
	pass := c.Query("pass")

	if user != "abdo" || pass != "abdo" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized",
		})

		return

	}

	c.Next()
}
