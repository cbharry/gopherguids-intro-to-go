package module03

import "fmt"

type Movie struct {
	Length  int
	Name    string
	plays   int
	rating  []float32
	viewers int
}

func (m *Movie) Rate(rating float32) error {
	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}
	m.rating = append(m.rating, rating)
	return nil
}

func (m *Movie) Play(viewers int) {
	m.plays++
	m.viewers += viewers
}

func (m Movie) Viewers() int {
	return m.viewers
}

func (m Movie) Plays() int {
	return m.plays
}

func (m Movie) Rating() float64 {
	var ratingTotal float32
	for _, r := range m.rating {
		ratingTotal += r
	}
	return float64(ratingTotal / float32(m.plays))
}

func (m Movie) String() string {
	return fmt.Sprintf("%s (%dm) %f%%", m.Name, m.Length, m.Rating())
}

type CritiqueFn func(m *Movie) (float32, error)

type Theatre struct{}

func (t *Theatre) Play(viewers int, movies ...*Movie) error {
	if len(movies) == 0 {
		return fmt.Errorf("no movies to play")
	}
	for _, m := range movies {
		m.Play(viewers)
	}
	return nil
}

func (t *Theatre) Critique(movies []*Movie, critique CritiqueFn) error {
	for _, m := range movies {
		m.Play(1)
		c, err := critique(m)
		if err != nil {
			return err
		}
		if err := m.Rate(c); err != nil {
			return err
		}
	}
	return nil
}
