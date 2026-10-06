package main

import (
	"context"
	"testing"
)

func TestHandleRequestReturnsNonEmptyMovieList(t *testing.T) {
	result, err := handleRequest(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) == 0 {
		t.Fatal("expected a non-empty list of movies")
	}
}
