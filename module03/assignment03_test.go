package module03

import (
	"fmt"
	"testing"
)

func TestRateAndPlay(t *testing.T) {
	movie := Movie{Name: "Lord of the Rings"}
	if err := movie.Rate(99); err == nil {
		t.Error("Shouldn't be allowed to rate without playing")
	}

	movie.Play(5)
	if err := movie.Rate(99); err != nil {
		t.Error("Couldn't rate")
	}

	if movie.Viewers() != 5 {
		t.Error("Viewers should be 5")
	}
	if movie.Plays() != 1 {
		t.Error("Plays should be 1")
	}
}

func TestRating(t *testing.T) {
	movie := Movie{Name: "Lord of the Rings"}
	movie.Play(1)
	movie.Play(1)
	movie.Rate(99)
	movie.Rate(95)
	if rating := movie.Rating(); rating != 97 {
		t.Errorf("Rating should be 97 got %f", rating)
	}
}

func TestString(t *testing.T) {
	movie := Movie{Name: "Lord of the Rings", Length: 178}
	movie.Play(1)
	movie.Rate(99)
	if fmt.Sprint(movie) != "Lord of the Rings (178m) 99.000000%" {
		t.Error("String function wrong")
	}
}

func TestTheatrePlay(t *testing.T) {
	theatre := Theatre{}
	lotr := Movie{Name: "LotR"}
	dh := Movie{Name: "Die Hard"}

	if err := theatre.Play(69); err == nil {
		t.Error("Theatre shouldn't be allowed to play without movies")
	}

	if err := theatre.Play(69, &lotr, &dh); err != nil {
		t.Error("Problem with playing theatre")
	}

	if lotr.Viewers() != 69 && dh.Viewers() != 69 {
		t.Error("Number of viewers wrong")
	}

	if lotr.Plays() != 1 && dh.Plays() != 1 {
		t.Error("Number of plays wrong")
	}
}

func TestTheatreCritique(t *testing.T) {
	theatre := Theatre{}
	movies := []*Movie{
		{Name: "LotR"},
		{Name: "Die Hard"},
	}
	c := CritiqueFn(func(m *Movie) (float32, error) {
		return 99.9, nil
	})
	theatre.Critique(movies, c)
	if movies[0].Rating() < 99.9 || movies[0].Rating() < 99.9 {
		t.Error("Critique failed")
	}
}
