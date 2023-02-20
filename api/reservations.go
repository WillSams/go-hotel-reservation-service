package api

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

type Reservation struct {
	ID           string  `db:"id"`
	RoomID       string  `db:"room_id"`
	CheckinDate  string  `db:"checkin_date"`
	CheckoutDate string  `db:"checkout_date"`
	TotalCharge  float64 `db:"total_charge"`
}

var reservationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Reservation",
	Fields: graphql.Fields{
		"Id":           &graphql.Field{Type: graphql.String},
		"RoomId":       &graphql.Field{Type: graphql.String},
		"CheckinDate":  &graphql.Field{Type: graphql.String},
		"CheckoutDate": &graphql.Field{Type: graphql.String},
		"TotalCharge":  &graphql.Field{Type: graphql.Float},
	},
})

var ReservationInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ReservationInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"RoomID": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"CheckinDate": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"CheckoutDate": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"TotalCharge": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
	},
})

func GetReservation(db *sqlx.DB) func(params graphql.ResolveParams) (interface{}, error) {
	return func(params graphql.ResolveParams) (interface{}, error) {
		roomId := params.Args["roomId"].(string)
		checkinDate := params.Args["roomId"].(string)
		checkoutDate := params.Args["roomId"].(string)
		var reservation Reservation
		err := db.Get(&reservation, "select id, room_id, checkin_date, checkout_date, total_charge from reservations room_id = $1 and checkin_date = $2 and checkout_date = $3 order by 1", roomId, checkinDate, checkoutDate)
		if err != nil {
			return nil, err
		}
		return reservation, nil
	}
}

func GetAllReservations(db *sqlx.DB) func(params graphql.ResolveParams) (interface{}, error) {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var reservations []Reservation
		err := db.Select(&reservations, "select id, room_id, checkin_date, checkout_date, total_charge from reservations order by 1")
		if err != nil {
			return nil, err
		}
		return reservations, nil
	}
}

func isRoomAvailable(db *sqlx.DB, roomID string, checkinDate string, checkoutDate string) (bool, error) {
	var count int

	err := db.Get(&count, `
		SELECT COUNT(*) 
		FROM reservations 
		WHERE room_id = $1 
		AND (
			(checkin_date >= $2 AND checkin_date < $3) OR 
			(checkout_date > $2 AND checkout_date <= $3) OR 
			(checkin_date <= $2 AND checkout_date >= $3)
		)
    `, roomID, checkinDate, checkoutDate)

	if err != nil {
		return false, nil
	}

	return count == 0, nil
}

func CreateReservation(db *sqlx.DB) func(graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		roomID := p.Args["roomId"].(string)

		// If there are any overlapping reservations, return an error
		available, err := isRoomAvailable(db, roomID, p.Args["checkinDate"].(string), p.Args["checkoutDate"].(string))
		if !available {
			return nil, fmt.Errorf("reservation dates overlap with an existing reservation")
		}
		checkinDate := p.Args["checkinDate"].(string)
		checkoutDate := p.Args["checkoutDate"].(string)
		totalCharge := p.Args["totalCharge"].(float64)

		reservation := Reservation{
			RoomID:       roomID,
			CheckinDate:  checkinDate,
			CheckoutDate: checkoutDate,
			TotalCharge:  totalCharge,
		}

		_, err = db.NamedExec(`
            INSERT INTO reservations (room_id, checkin_date, checkout_date, total_charge)
            VALUES (:room_id, :checkin_date, :checkout_date, :total_charge)
        `, reservation)
		if err != nil {
			return nil, err
		}

		return reservation, nil
	}
}
