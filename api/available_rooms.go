package api

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

type Room struct {
	ID           string  `db:"id"`
	NumBeds      int     `db:"num_beds"`
	AllowSmoking bool    `db:"allow_smoking"`
	DailyRate    float64 `db:"daily_rate"`
	CleaningFee  float64 `db:"cleaning_fee"`
	TotalCharge  float64 `db:"total_charge"`
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

func daysBetween(startDateStr string, endDateStr string) int {
	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)
	days := int(endDate.Sub(startDate).Hours() / 24)
	return days
}

func GetAvailableRooms(db *sqlx.DB) func(params graphql.ResolveParams) (interface{}, error) {
	return func(params graphql.ResolveParams) (interface{}, error) {
		startDate, _ := params.Args["startDate"].(string)
		numBeds, _ := params.Args["numBeds"].(int)
		allowSmoking, _ := params.Args["allowSmoking"].(bool)
		endDate, _ := params.Args["endDate"].(string)

		numDays := daysBetween(startDate, endDate)

		// Query the database for available rooms
		query := fmt.Sprintf(`select distinct ro.id, ro.num_beds, ro.allow_smoking, ro.daily_rate, ro.cleaning_fee,
				((ro.daily_rate * %d + ro.cleaning_fee)) as total_charge
			from rooms as ro
			where ro.allow_smoking = %t
				and ro.num_beds >= %d
				and ro.id not in (
				select room_id 
				from reservations 
				where '%s' between checkin_date and checkout_date
					
				and '%s'  between checkin_date and checkout_date
				)
			order by 6, 2`, numDays, allowSmoking, numBeds, endDate, startDate)

		var rooms []Room
		err := db.Select(&rooms, query)
		if err != nil {
			return nil, err
		}

		return rooms, nil
	}
}
