package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/trakfy/backend/db"
	"github.com/trakfy/backend/routes"
)

func main() {
	godotenv.Load()

	db.RunMigrations()

	err := db.SetupDB()
	if err != nil {
		return
	}

	router := gin.Default()
	routes.SetupRoutes(router)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8000"
	}
	router.Run(port)
}
