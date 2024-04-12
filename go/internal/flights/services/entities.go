package services

import "airports/internal/flights/repositories"

type Flight struct {
	ID                 uint64 `db:"flight_id" json:"id"`
	Number             string `db:"flight_no" json:"number"`
	Status             string `db:"status" json:"status"`
	AircraftCode       string `db:"aircraft_code" json:"aircraft_code"`
	ScheduledDeparture string `db:"scheduled_departure" json:"scheduled_departure"`
	ScheduledArrival   string `db:"scheduled_arrival" json:"scheduled_arrival"`
}

func BuildFlightAggregate(data *repositories.GetFlightOutput) (*FlightAggregate, error) {

	return &FlightAggregate{}, nil
}

func BuildFlights(data *repositories.GetFlightsOutput) ([]Flight, error) {
	flights := []Flight{}

	return flights, nil
}
