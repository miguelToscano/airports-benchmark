package repositories

import (
	"airports/internal/database"
)

type Flight struct {
	ID                 uint64 `db:"flight_id"`
	Number             string `db:"flight_no"`
	Status             string `db:"status"`
	AircraftCode       string `db:"aircraft_code"`
	ScheduledDeparture string `db:"scheduled_departure"`
	ScheduledArrival   string `db:"scheduled_arrival"`
}

type GetFlightsOutput = []Flight

func (fr *FlightsRepository) GetFlights() (*GetFlightsOutput, error) {
	flights := GetFlightsOutput{}

	if err := database.DB.Select(&flights, "SELECT flight_id, flight_no, status, aircraft_code, scheduled_departure, scheduled_arrival FROM flights ORDER BY flight_id ASC LIMIT 100"); err != nil {
		return nil, err
	}

	return &flights, nil
}
