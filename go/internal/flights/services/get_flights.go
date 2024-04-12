package services

import (
	"airports/internal/flights/domain"
	"fmt"
)

func (fs *FlightsService) GetFlights() ([]domain.Flight, error) {
	data, err := fs.flightsRepository.GetFlights()

	if err != nil {
		return nil, fmt.Errorf("GetFlights error: %w", err)
	}

	flights, err := domain.BuildFlights(data)

	if err != nil {
		return nil, fmt.Errorf("BuildFlights error: %w", err)
	}

	return flights, nil
}
