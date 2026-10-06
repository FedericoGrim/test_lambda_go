package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/FedericoGrim/test_lambda_go/movies"
)

func handleRequest(ctx context.Context) ([]movies.Movie, error) {
	return movies.All(), nil
}

func main() {
	lambda.Start(handleRequest)
}
