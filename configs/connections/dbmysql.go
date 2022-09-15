package connections

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDb() *gorm.DB {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	envPath := filepath.Dir(wd) + "\\.env"
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file1")
	}
	dsn := os.Getenv("MYSQL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
