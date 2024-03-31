package services

import "airports/internal/flights/domain"

func (fs *FlightsService) GetFlight(id uint64) (*domain.FlightAggregate, error) {
	flight, err := fs.flightsRepository.GetFlight(id)

	if err != nil {
		return nil, err
	}

	return flight, nil
}
