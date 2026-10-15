package grade

import (
	"errors"
	"testing"
)

func TestLetter(t *testing.T) {
	cases := []struct {
		name  string
		score int
		want  string
	}{
		{"top", 100, "A"}, {"a floor", 90, "A"}, {"below a", 89, "B"},
		{"b floor", 80, "B"}, {"below b", 79, "C"}, {"c floor", 70, "C"},
		{"below c", 69, "D"}, {"d floor", 60, "D"}, {"below d", 59, "F"}, {"zero", 0, "F"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Letter(c.score)
			if err != nil || got != c.want {
				t.Fatalf("Letter(%d) = %q, %v; want %q", c.score, got, err, c.want)
			}
		})
	}
}

func TestLetterRange(t *testing.T) {
	for _, s := range []int{-1, 101} {
		if _, err := Letter(s); !errors.Is(err, ErrRange) {
			t.Errorf("Letter(%d) err = %v", s, err)
		}
	}
}
