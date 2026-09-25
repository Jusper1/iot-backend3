package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"iot-backend/internal/config"
	"iot-backend/internal/routes"
)

func main() {


	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env tidak ditemukan")
	}

	// Connect database
	config.ConnectDatabase()

	// Gin
	r := gin.Default()

	// CORS

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		AllowCredentials: true,
	}))

	// ROUTES
	routes.SetupRoutes(r, config.DB)

	// SERVER

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server berjalan di port:", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}