package repositories

import (
	"airports/internal/database"
	"fmt"
)

type GetFlightOutput struct {
	ID                 uint64 `db:"flight_id" json:"id"`
	Number             string `db:"flight_no" json:"number"`
	Status             string `db:"status" json:"status"`
	AircraftCode       string `db:"aircraft_code" json:"aircraft_code"`
	ScheduledDeparture string `db:"scheduled_departure" json:"scheduled_departure"`
	ScheduledArrival   string `db:"scheduled_arrival" json:"scheduled_arrival"`
}

func (fr *FlightsRepository) GetFlight(id uint64) (*GetFlightOutput, error) {
	flight := GetFlightOutput{}

	query := `
		SELECT
			flight_id,
			flight_no,
			status,
			scheduled_departure,
			scheduled_arrival,
			JSONB_BUILD_OBJECT(
				'airport_code',
				a1.airport_code,
				'airport_name',
				a1.airport_name,
				'city',
				a1.city,
				'timezone',
				a1.timezone
			) AS arrival_airport,
			JSONB_BUILD_OBJECT(
				'airport_code',
				a2.airport_code,
				'airport_name',
				a2.airport_name,
				'city',
				a2.city,
				'timezone',
				a2.timezone
			) AS departure_airport,
			JSONB_BUILD_OBJECT('aircraft_code',
			a3.aircraft_code,
			'model',
			a3.model,
			'range',
			a3.range) AS aircraft
		FROM
			flights f
		LEFT JOIN airports a1 ON
			a1.airport_code = f.arrival_airport
		LEFT JOIN airports a2 ON
			a2.airport_code = f.departure_airport
		LEFT JOIN aircrafts a3 USING(aircraft_code)
		WHERE flight_id=$1
	`

	if err := database.DB.Get(&flight, query, id); err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &flight, nil
}
