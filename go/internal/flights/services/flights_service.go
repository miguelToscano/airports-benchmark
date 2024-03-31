package services

import "airports/internal/flights/domain"

type FlightsRepository interface {
	GetFlights() ([]domain.Flight, error)
	GetFlight(id uint64) (*domain.FlightAggregate, error)
}

type FlightsService struct {
	flightsRepository FlightsRepository
}

func NewFlightsService(flightsRepository FlightsRepository) *FlightsService {
	return &FlightsService{
		flightsRepository: flightsRepository,
	}
}
