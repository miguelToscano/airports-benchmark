package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"airports/internal/database"
	flightsController "airports/internal/flights/controllers"
	"airports/internal/flights/repositories"
	"airports/internal/flights/services"
)

func setupEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func setupDatabase() {
	err := database.Connect()

	if err != nil {
		panic("Error connecting database")
	}
}

func start() {
	flightsRepository := repositories.NewFlightsRepository()
	flightsService := services.NewFlightsService(flightsRepository)
	flightsController := flightsController.NewFlightsController(flightsService)

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/v1/flights", flightsController.GetFlights)
	app.Get("/v1/flights/:id", flightsController.GetFlight)

	app.Listen(":3000")
}

func main() {
	setupEnvironmentVariables()
	setupDatabase()
	start()
}
