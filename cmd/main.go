package main

import (
	"fmt"
	"log"

	"github.com/almazatun/lien-court/pkg/common/helper"
	"github.com/almazatun/lien-court/pkg/database"
	routes "github.com/almazatun/lien-court/pkg/http"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {
	loadENVs()
	err := database.ConnectDB()

	if err != nil {
		fmt.Println("🔴 Database connection error")
		log.Fatal(err)
	}

	app := fiber.New()

	routes.PublicRoutes(app)

	log.Fatal(app.Listen(helper.GetEnvVar("PORT")))
}

func loadENVs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("🔴 Error loading .env variables")
	}
}
