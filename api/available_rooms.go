package api

import (
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

type Room struct {
	ID           string  `db:"id"`
	NumBeds      int     `db:"num_beds"`
	AllowSmoking bool    `db:"allow_smoking"`
	DailyRate    float64 `db:"daily_rate"`
	CleaningFee  float64 `db:"cleaning_fee"`
}

var roomType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Room",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.String,
		},
		"NumBeds": &graphql.Field{
			Type: graphql.Int,
		},
		"AllowSmoking": &graphql.Field{
			Type: graphql.Boolean,
		},
		"DailyRate": &graphql.Field{
			Type: graphql.Float,
		},
		"CleaningFee": &graphql.Field{
			Type: graphql.Float,
		},
		"TotalCharge": &graphql.Field{
			Type: graphql.Float,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				room, _ := params.Source.(*Room)
				return room.DailyRate + room.CleaningFee, nil
			},
		},
	},
})

func GetAvailableRooms(db *sqlx.DB) func(params graphql.ResolveParams) (interface{}, error) {
	return func(params graphql.ResolveParams) (interface{}, error) {
		startDate, _ := params.Args["startDate"].(string)
		numBeds, _ := params.Args["numBeds"].(int)
		allowSmoking, _ := params.Args["allowSmoking"].(bool)
		endDate, _ := params.Args["endDate"].(string)

		// Query the database for available rooms
		query := `select distinct ro.id, ro.num_beds, ro.allow_smoking, ro.daily_rate, ro.cleaning_fee,
				((ro.daily_rate * $2 + ro.cleaning_fee) as total_charge
			from rooms as ro
			where ro.allow_smoking = $3
				and ro.num_beds >= $2
				and ro.id not in (
				select room_id 
				from reservations 
				where '$4' between checkin_date and checkout_date
					and '$1'  between checkin_date and checkout_date
				)
			order by 6, 2`

		var rooms []Room
		err := db.Select(&rooms, query, numBeds, allowSmoking, endDate, startDate)
		if err != nil {
			return nil, err
		}

		return rooms, nil
	}
}
