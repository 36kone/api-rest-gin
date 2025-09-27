package routes

import (
	"github.com/36kone/api-go/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/users", controllers.ReadAllUsers)
	r.POST("/users", controllers.CreateUser)
	err := r.Run()
	if err != nil {
		return
	}
}
