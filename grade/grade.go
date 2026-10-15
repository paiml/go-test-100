// Package grade turns a score into a letter. Its boundaries are the point:
// each one needs a test on both sides.
package grade

import "errors"

// ErrRange is returned for a score outside 0..100.
var ErrRange = errors.New("score out of range 0..100")

var floors = []struct {
	min    int
	letter string
}{{90, "A"}, {80, "B"}, {70, "C"}, {60, "D"}}

// Letter maps 0..100 to A, B, C, D or F.
func Letter(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", ErrRange
	}
	for _, f := range floors {
		if score >= f.min {
			return f.letter, nil
		}
	}
	return "F", nil
}
