package routes

import (
	"../services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/users", services.CreateUser)
	r.GET("/user/:id", services.GetUser)
	r.PATCH("/user/:id", services.UpdateUser)
}
