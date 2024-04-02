package main

import (
	"dizeto-backend/app/repository"
	"dizeto-backend/app/router"
	"dizeto-backend/app/service"
	"dizeto-backend/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize database connection and perform auto migrate
	db, err := config.InitDB()
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepo)

	// Initialize router
	r := gin.Default()
	router.SetupRouter(r, userService)

	// Run the application
	r.Run(":8080")
}
