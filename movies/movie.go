package movies

import "strings"

type Movie struct {
	Title    string
	Director string
	Genre    string
	Actors   []string
}

var catalog = []Movie{
	{
		Title:    "The Godfather",
		Director: "Francis Ford Coppola",
		Genre:    "Crime",
		Actors:   []string{"Marlon Brando", "Al Pacino", "James Caan"},
	},
	{
		Title:    "Inception",
		Director: "Christopher Nolan",
		Genre:    "Sci-Fi",
		Actors:   []string{"Leonardo DiCaprio", "Joseph Gordon-Levitt", "Elliot Page"},
	},
	{
		Title:    "Parasite",
		Director: "Bong Joon-ho",
		Genre:    "Thriller",
		Actors:   []string{"Song Kang-ho", "Lee Sun-kyun", "Cho Yeo-jeong"},
	},
	{
		Title:    "The Grand Budapest Hotel",
		Director: "Wes Anderson",
		Genre:    "Comedy",
		Actors:   []string{"Ralph Fiennes", "Tony Revolori", "Saoirse Ronan"},
	},
	{
		Title:    "Spirited Away",
		Director: "Hayao Miyazaki",
		Genre:    "Animation",
		Actors:   []string{"Rumi Hiiragi", "Miyu Irino"},
	},
}

func All() []Movie {
	return catalog
}

func ByTitle(title string) (Movie, bool) {
	for _, movie := range catalog {
		if strings.EqualFold(movie.Title, title) {
			return movie, true
		}
	}
	return Movie{}, false
}

func ByGenre(genre string) []Movie {
	return filter(func(movie Movie) bool {
		return strings.EqualFold(movie.Genre, genre)
	})
}

func ByDirector(director string) []Movie {
	return filter(func(movie Movie) bool {
		return strings.EqualFold(movie.Director, director)
	})
}

func filter(matches func(Movie) bool) []Movie {
	var results []Movie
	for _, movie := range catalog {
		if matches(movie) {
			results = append(results, movie)
		}
	}
	return results
}
