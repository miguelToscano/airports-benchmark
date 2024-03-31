package services

import "airports/internal/flights/domain"

func (fs *FlightsService) GetFlights() ([]domain.Flight, error) {
	flights, err := fs.flightsRepository.GetFlights()

	if err != nil {
		return nil, err
	}

	return flights, nil
}
