package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/FedericoGrim/test_lambda_go/movies"
)

type Request struct {
	Genre string `json:"genre"`
}

func handleRequest(ctx context.Context, request Request) ([]movies.Movie, error) {
	return movies.ByGenre(request.Genre), nil
}

func main() {
	lambda.Start(handleRequest)
}
