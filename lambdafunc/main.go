package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/willsams/go-hotel-reservation-service/api"
)

func main() {
	lambda.Start(api.GraphQlApiHandler)
	//api.DebubGraphQlApiHandler()   // for local testing
}
