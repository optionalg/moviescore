package getratings

import (
	"strings"
)

type ImdbMovie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Genre      string
	Director   string
	Actors     string
	Plot       string
	Awards     string
	Poster     string
	Metascore  string
	ImdbRating string `json:"imdbRating"`
}

func GetImdbRatings(mname string) ImdbMovie {
	movieName := strings.Replace(mname, " ", "+", 9)
	movieInfo := new(ImdbMovie)
	GetJSON("http://www.omdbapi.com/?t="+movieName+"&plot=full", movieInfo)
	return *movieInfo
}
