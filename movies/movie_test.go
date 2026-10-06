package movies

import (
	"strings"
	"testing"
)

func TestAllReturnsNonEmptyCatalog(t *testing.T) {
	all := All()

	if len(all) == 0 {
		t.Fatal("expected a non-empty catalog")
	}
}

func TestByTitleFindsExistingMovieCaseInsensitively(t *testing.T) {
	movie, found := ByTitle("inception")

	if !found {
		t.Fatal("expected to find a movie titled Inception")
	}
	if movie.Title != "Inception" {
		t.Errorf("got title %q, want %q", movie.Title, "Inception")
	}
}

func TestByTitleReturnsFalseWhenNotFound(t *testing.T) {
	_, found := ByTitle("Nonexistent Movie")

	if found {
		t.Fatal("expected not to find a nonexistent movie")
	}
}

func TestByGenreFiltersCaseInsensitively(t *testing.T) {
	results := ByGenre("sci-fi")

	if len(results) == 0 {
		t.Fatal("expected at least one Sci-Fi movie")
	}
	for _, movie := range results {
		if !strings.EqualFold(movie.Genre, "sci-fi") {
			t.Errorf("got genre %q, want Sci-Fi", movie.Genre)
		}
	}
}

func TestByGenreReturnsEmptyWhenNoMatch(t *testing.T) {
	results := ByGenre("Nonexistent Genre")

	if len(results) != 0 {
		t.Errorf("got %d results, want 0", len(results))
	}
}

func TestByDirectorFiltersCaseInsensitively(t *testing.T) {
	results := ByDirector("christopher nolan")

	if len(results) == 0 {
		t.Fatal("expected at least one movie directed by Christopher Nolan")
	}
	for _, movie := range results {
		if !strings.EqualFold(movie.Director, "christopher nolan") {
			t.Errorf("got director %q, want Christopher Nolan", movie.Director)
		}
	}
}

func TestByDirectorReturnsEmptyWhenNoMatch(t *testing.T) {
	results := ByDirector("Nonexistent Director")

	if len(results) != 0 {
		t.Errorf("got %d results, want 0", len(results))
	}
}
