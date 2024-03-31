package controllers

import "airports/internal/flights/domain"

type FlightsService interface {
	GetFlights() ([]domain.Flight, error)
	GetFlight(id uint64) (*domain.FlightAggregate, error)
}

type FlightsController struct {
	flightsService FlightsService
}

func NewFlightsController(flightsService FlightsService) *FlightsController {
	return &FlightsController{
		flightsService: flightsService,
	}
}
