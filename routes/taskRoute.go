package routes

import (
	"github.com/gin-gonic/gin"
	"test-v2/controllers"
)

func TaskRoutes(r *gin.Engine) {
	router := r.Group("t")
	router.GET("/tasks", controllers.FindTasks)
	router.POST("/task", controllers.CreateTask)
	router.GET("task/:id", controllers.FindTask)
	router.PATCH("tasks/:id", controllers.UpdateTask)
	router.DELETE("tasks/:id", controllers.DeleteTask)

}
