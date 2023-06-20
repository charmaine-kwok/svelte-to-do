package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupDoneRoutes(grp *gin.RouterGroup) {
	grp.POST("", controllers.PostDone)
	grp.GET("/dones", controllers.GetAllDones)
	grp.DELETE("/:id", controllers.DeleteDone)
}
