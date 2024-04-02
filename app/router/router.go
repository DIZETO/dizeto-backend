package router

import (
	"dizeto-backend/app/controller"
	"dizeto-backend/app/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userService service.UserService) {
	userController := controller.NewUserController(userService)

	// Routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
	}
}
