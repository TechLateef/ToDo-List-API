package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/techlateef/tech-lateef-gol/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	DbName := os.Getenv("DB_DATABASE")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.ToDoModels{})
	return db
}

func CloseDatabaseC(db *gorm.DB) {
	dbMsql, err := db.DB()
	if err != nil {
		panic("Failed to close db")

	}
	dbMsql.Close()
}
