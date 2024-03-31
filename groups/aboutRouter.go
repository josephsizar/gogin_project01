package groups

import (
	"goveri/controllers"

	"github.com/gin-gonic/gin"
)

func About(g *gin.RouterGroup) {
	g.GET("/", controllers.AboutIndex)
}
