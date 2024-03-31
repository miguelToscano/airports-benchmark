package repositories

import (
	"airports/internal/database"
	"airports/internal/flights/domain"
	"fmt"
)

func (fr *FlightsRepository) GetFlight(id uint64) (*domain.FlightAggregate, error) {
	flight := &domain.FlightAggregate{}

	if err := database.DB.Get(flight, `
	select
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
		) as arrival_airport,
		JSONB_BUILD_OBJECT(
			'airport_code',
			a2.airport_code,
			'airport_name',
			a2.airport_name,
			'city',
			a2.city,
			'timezone',
			a2.timezone
		) as departure_airport,
		jsonb_build_object('aircraft_code',
		a3.aircraft_code,
		'model',
		a3.model,
		'range',
		a3.range) as aircraft
	from
		flights f
	left join airports a1 on
		a1.airport_code = f.arrival_airport
	left join airports a2 on
		a2.airport_code = f.departure_airport
	left join aircrafts a3 on
	a3.aircraft_code = f.aircraft_code WHERE flight_id=$1`, id); err != nil {
		fmt.Print(err)
		return nil, err
	}

	return flight, nil
}
