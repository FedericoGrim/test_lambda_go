package main

import (
	"context"
	"testing"
)

func TestHandleRequestReturnsMoviesMatchingGenre(t *testing.T) {
	result, err := handleRequest(context.Background(), Request{Genre: "Sci-Fi"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) == 0 {
		t.Fatal("expected at least one Sci-Fi movie")
	}
}

func TestHandleRequestReturnsEmptyListWhenGenreNotFound(t *testing.T) {
	result, err := handleRequest(context.Background(), Request{Genre: "Nonexistent Genre"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("got %d results, want 0", len(result))
	}
}
