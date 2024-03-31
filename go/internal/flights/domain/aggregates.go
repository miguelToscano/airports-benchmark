package domain

import "encoding/json"

type FlightAggregate struct {
	ID                 uint64          `db:"flight_id" json:"id"`
	Number             string          `db:"flight_no" json:"number"`
	Status             string          `db:"status" json:"status"`
	ScheduledDeparture string          `db:"scheduled_departure" json:"scheduled_departure"`
	ScheduledArrival   string          `db:"scheduled_arrival" json:"scheduled_arrival"`
	ArrivalAiport      json.RawMessage `db:"arrival_airport" json:"arrival_airport"`
	DepartureAirport   json.RawMessage `db:"departure_airport" json:"departure_airport"`
	Aircraft           json.RawMessage `db:"aircraft" json:"aircraft"`
}
