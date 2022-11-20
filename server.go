package main

import (
	"github.com/techlateef/tech-lateef-gol/configs"
	"github.com/techlateef/tech-lateef-gol/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = configs.ConnectToDB()
)

func main() {
	defer configs.CloseDatabaseC(db)
	//run all routes

	routes.Routes()

}
