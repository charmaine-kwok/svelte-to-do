package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(grp *gin.RouterGroup) {
	grp.POST("", controllers.PostTodo)
	grp.GET("/todos", controllers.GetAllTodos)
	grp.DELETE("/:id", controllers.DeleteTodo)
	grp.PATCH("/:id", controllers.UpdateTodo)
}
