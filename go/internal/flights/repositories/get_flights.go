package repositories

import (
	"airports/internal/database"
	"airports/internal/flights/domain"
)

func (fr *FlightsRepository) GetFlights() ([]domain.Flight, error) {
	flights := []domain.Flight{}

	if err := database.DB.Select(&flights, "SELECT flight_id, flight_no, status, aircraft_code, scheduled_departure, scheduled_arrival FROM flights ORDER BY flight_id ASC LIMIT 100"); err != nil {
		return nil, err
	}

	return flights, nil
}
