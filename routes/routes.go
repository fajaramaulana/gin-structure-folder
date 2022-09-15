package routes

import (
	"go-structure-folder-clean/cmd/controllers"
	"go-structure-folder-clean/cmd/repositories"
	"go-structure-folder-clean/cmd/services"
	"go-structure-folder-clean/configs/connections"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	App := gin.Default()

	db := connections.ConnectionDb()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	envPath := filepath.Dir(wd) + "\\.env"
	err = godotenv.Load(envPath)
	if err != nil {
		// log.Fatal("Error loading .env file")
		log.Fatal(err.Error())
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
