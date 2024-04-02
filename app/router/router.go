package router

import (
	"dizeto-backend/app/controller"
	"dizeto-backend/app/repository"
	"dizeto-backend/app/service"
	"dizeto-backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	// Initialize repository
	userRepo := repository.NewUserRepository(db)
	aboutRepo := repository.NewAboutRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepo)
	aboutService := service.NewAboutService(aboutRepo)

	// Initialize controller
	userController := controller.NewUserController(userService)
	aboutController := controller.NewAboutController(aboutService)

	// Routes
	v1 := r.Group("/api/v1")
	{
		//user
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)

		//about
		v1.POST("/about", middleware.AuthorizationMiddleware(), aboutController.CreateAbout)
		v1.GET("/about", aboutController.GetAbout)

	}
}
