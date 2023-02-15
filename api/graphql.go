package api

import (
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

func AppSchema(db *sqlx.DB) (graphql.Schema, error) {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
		"availableRooms": &graphql.Field{
			Type: graphql.NewList(roomType),
			Args: graphql.FieldConfigArgument{
				"startDate": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"numBeds": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"allowSmoking": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
				"endDate": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: GetAvailableRooms(db),
		},
		"reservations": &graphql.Field{
			Type:    graphql.NewList(reservationType),
			Resolve: GetAllReservations(db),
		},
		"reservation": &graphql.Field{
			Type: reservationType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: GetReservation(db),
		},
	}}

	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: graphql.Fields{
		"createReservation": &graphql.Field{
			Type:        reservationType,
			Description: "Create a reservation",
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(ReservationInputType),
				},
			},
			Resolve: CreateReservation(db),
		},
	}}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}

	return graphql.NewSchema(schemaConfig)
}
