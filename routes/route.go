package routes

import(
	"test-v2/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r:= gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	}) 
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/task", controllers.CreateTask)
	r.GET("task/:id", controllers.FindTask)
	r.PATCH("tasks/:id", controllers.UpdateTask)
	r.DELETE("tasks/:id", controllers.DeleteTask)

	return r
}