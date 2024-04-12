package services

import (
	"airports/internal/flights/repositories"
)

type FlightsRepository interface {
	GetFlights() (*repositories.GetFlightsOutput, error)
	GetFlight(id uint64) (*repositories.GetFlightOutput, error)
}

type FlightsService struct {
	flightsRepository FlightsRepository
}

func NewFlightsService(flightsRepository FlightsRepository) *FlightsService {
	return &FlightsService{
		flightsRepository: flightsRepository,
	}
}
