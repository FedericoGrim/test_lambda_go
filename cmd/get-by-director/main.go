package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/FedericoGrim/test_lambda_go/movies"
)

type Request struct {
	Director string `json:"director"`
}

func handleRequest(ctx context.Context, request Request) ([]movies.Movie, error) {
	return movies.ByDirector(request.Director), nil
}

func main() {
	lambda.Start(handleRequest)
}
