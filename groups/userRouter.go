package groups

import (
	"goveri/controllers"

	"github.com/gin-gonic/gin"
)

func User(g *gin.RouterGroup) {

	g.GET("/", controllers.UserIndex)

	g.GET("/all", controllers.UserAll)

	g.GET("/:id", controllers.UserById)

}
