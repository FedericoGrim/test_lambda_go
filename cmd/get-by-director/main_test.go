package main

import (
	"context"
	"testing"
)

func TestHandleRequestReturnsMoviesMatchingDirector(t *testing.T) {
	result, err := handleRequest(context.Background(), Request{Director: "Christopher Nolan"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) == 0 {
		t.Fatal("expected at least one movie directed by Christopher Nolan")
	}
}

func TestHandleRequestReturnsEmptyListWhenDirectorNotFound(t *testing.T) {
	result, err := handleRequest(context.Background(), Request{Director: "Nonexistent Director"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("got %d results, want 0", len(result))
	}
}
