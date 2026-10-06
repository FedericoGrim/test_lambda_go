package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/FedericoGrim/test_lambda_go/movies"
)

type Request struct {
	Title string `json:"title"`
}

func handleRequest(ctx context.Context, request Request) (movies.Movie, error) {
	movie, found := movies.ByTitle(request.Title)
	if !found {
		return movies.Movie{}, fmt.Errorf("movie not found: %s", request.Title)
	}
	return movie, nil
}

func main() {
	lambda.Start(handleRequest)
}
