package domain

type Flight struct {
	ID                 uint64 `db:"flight_id" json:"id"`
	Number             string `db:"flight_no" json:"number"`
	Status             string `db:"status" json:"status"`
	AircraftCode       string `db:"aircraft_code" json:"aircraft_code"`
	ScheduledDeparture string `db:"scheduled_departure" json:"scheduled_departure"`
	ScheduledArrival   string `db:"scheduled_arrival" json:"scheduled_arrival"`
}

type Airport struct {
	AirportCode string `db:"airport_code" json:"airport_code"`
	AirportName string `db:"airport_name" json:"airport_name"`
	City        string `db:"city" json:"city"`
	Timezone    string `db:"timezone" json:"timezone"`
}

type Aircraft struct {
	AircraftCode string `db:"aircraft_code" json:"aircraft_code"`
	Model        string `db:"model" json:"model"`
	Range        int    `db:"range"  json:"range"`
}
