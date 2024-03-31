package controllers

import (
	"airports/internal/flights/domain"

	"github.com/gofiber/fiber/v2"
)

type GetFlightsDTO struct {
	Total   int             `json:"total"`
	Flights []domain.Flight `json:"flights"`
}

func (fc *FlightsController) GetFlights(c *fiber.Ctx) error {
	flights, err := fc.flightsService.GetFlights()

	if err != nil {
		return c.Status(500).JSON(map[string]string{
			"message": "error",
		})
	}

	response := GetFlightsDTO{
		Flights: flights,
		Total:   len(flights),
	}

	return c.JSON(response)
}
