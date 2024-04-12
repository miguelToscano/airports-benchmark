package services

import (
	"airports/internal/errors"
	"airports/internal/flights/domain"
	"fmt"
)

func (fs *FlightsService) GetFlight(id uint64) (*domain.FlightAggregate, error) {
	data, err := fs.flightsRepository.GetFlight(id)

	if err != nil {
		return nil, fmt.Errorf("GetFlightError: %w", err)
	}

	flight, err := domain.BuildFlightAggregate(data)

	if err != nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("Flight with id: %d not found", id))
	}

	if flight == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("Flight with id: %d not found", id))
	}

	return flight, nil
}
