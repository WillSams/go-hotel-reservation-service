package api

import (
	"github.com/google/uuid"
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

func CreateReservation(db *sqlx.DB) func(params graphql.ResolveParams) (interface{}, error) {
	return func(params graphql.ResolveParams) (interface{}, error) {
		roomID := params.Args["roomId"].(string)
		checkinDate := params.Args["checkinDate"].(string)
		checkoutDate := params.Args["checkoutDate"].(string)
		totalCharge := params.Args["totalCharge"].(float64)

		var reservation Reservation
		reservation.ID = uuid.New().String()
		reservation.RoomID = roomID
		reservation.CheckinDate = checkinDate
		reservation.CheckoutDate = checkoutDate
		reservation.TotalCharge = totalCharge

		_, err := db.NamedExec(`
			INSERT INTO reservations (id, room_id, checkin_date, checkout_date, total_charge)
			VALUES (:id, :room_id, :checkin_date, :checkout_date, :total_charge)
		`, reservation)
		if err != nil {
			return nil, err
		}

		return reservation, nil
	}
}
