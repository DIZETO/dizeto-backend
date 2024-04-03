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
	highlightRepo := repository.NewHighlightPortofolio(db)
	pricingRepo := repository.NewPricingRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepo)
	aboutService := service.NewAboutService(aboutRepo)
	highlightService := service.NewHighlightService(highlightRepo)
	pricingService := service.NewPricingService(pricingRepo)

	// Initialize controller
	userController := controller.NewUserController(userService)
	aboutController := controller.NewAboutController(aboutService)
	highlightController := controller.NewHighlightController(highlightService)
	pricingController := controller.NewPricingController(pricingService)

	// Routes
	v1 := r.Group("/api/v1")
	{
		//user
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)

		//about
		v1.POST("/about", middleware.AuthorizationMiddleware(), aboutController.CreateAbout)
		v1.GET("/about", aboutController.GetAllAbout)
		v1.PUT("/about/:id", middleware.AuthorizationMiddleware(), aboutController.UpdateAbout)

		//highlight
		v1.POST("/highlight_portofolio", middleware.AuthorizationMiddleware(), highlightController.CreateHighlight)
		v1.GET("/highlight_portofolio", highlightController.GetAllHighlight)
		v1.GET("/highlight_portofolio/:id", highlightController.GetHighlightByID)
		v1.PUT("/highlight_portofolio/:id", middleware.AuthorizationMiddleware(), highlightController.UpdateHighlight)

		//pricing
		v1.POST("/pricing", middleware.AuthorizationMiddleware(), pricingController.CreatePricing)
		v1.GET("/pricing", pricingController.GetAllPricing)
		v1.PUT("/pricing/:id", middleware.AuthorizationMiddleware(), pricingController.UpdatePricing)
	}
}
