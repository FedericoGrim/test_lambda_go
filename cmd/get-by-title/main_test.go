package main

import (
	"context"
	"testing"
)

func TestHandleRequestFindsExistingMovie(t *testing.T) {
	result, err := handleRequest(context.Background(), Request{Title: "Inception"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Title != "Inception" {
		t.Errorf("got title %q, want %q", result.Title, "Inception")
	}
}

func TestHandleRequestReturnsErrorWhenMovieNotFound(t *testing.T) {
	_, err := handleRequest(context.Background(), Request{Title: "Nonexistent Movie"})

	if err == nil {
		t.Fatal("expected an error for a nonexistent movie")
	}
}
