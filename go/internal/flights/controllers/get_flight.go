package controllers

import (
	"airports/internal/flights/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type GetFlightResponse struct {
	Flight *domain.FlightAggregate `json:"flight"`
}

func (fc *FlightsController) GetFlight(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return err
	}

	flight, err := fc.flightsService.GetFlight(id)

	if err != nil {
		return c.Status(500).JSON(map[string]string{
			"message": "error",
		})
	}

	response := GetFlightResponse{
		Flight: flight,
	}

	return c.JSON(response)
}
