package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Define the GraphQL schema
var schema, _ = AppSchema(dbConnect())

// Define the available rooms handler
func GraphQlApiHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db := dbConnect()
	defer db.Close()

	queryParameters := make(map[string]interface{})
	for key, value := range request.QueryStringParameters {
		queryParameters[key] = value
	}

	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  request.Body,
		Context:        ctx,
		VariableValues: queryParameters,
	})

	return buildAPIGatewayResponse(200, result)
}

func dbConnect() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func buildAPIGatewayResponse(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	responseBody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            string(responseBody),
		Headers: map[string]string{
			"Content-Type":             "application/json",
			"X-YOURCOMPANY-Func-Reply": "graphql-api-handler",
		},
	}, nil
}
