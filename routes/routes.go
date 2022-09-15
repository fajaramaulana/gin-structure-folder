package routes

import (
	"go-structure-folder-clean/cmd/controllers"
	"go-structure-folder-clean/cmd/repositories"
	"go-structure-folder-clean/cmd/services"
	"go-structure-folder-clean/configs/connections"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	App := gin.Default()

	db := connections.ConnectionDb()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	modeGin := os.Getenv("GIN_MODE")

	// auth service
	authService := services.AuthNewService()

	// user
	userRepository := repositories.NewRepository(db)
	userService := services.UserNewService(userRepository)
	userHandler := controllers.NewUserHandler(userService, authService)

	gin.SetMode(modeGin)

	v1 := App.Group("/api/v1")

	userHandler.AuthRoutes(v1)

	App.Run(os.Getenv("PORT"))

}
