package Routes

import (
	"vahidid/20-project-go/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	taskRoutes := r.Group("tasks")
	{
		taskRoutes.GET("", Controllers.ListTask)
		taskRoutes.POST("", Controllers.CreateTask)
		taskRoutes.GET("/:id", Controllers.GetTask)
		taskRoutes.PUT("/:id", Controllers.UpdateTask)
		taskRoutes.DELETE("/:id", Controllers.DeleteTask)
	}

	return r
}
