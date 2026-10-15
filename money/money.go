// Package money parses and formats whole cents. Parse and Format must
// round-trip, which the fuzz test checks.
package money

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// maxWhole keeps dollars*100 far from int overflow.
const maxWhole = 12

// ErrFormat is returned for text that is not dollars.cents.
var ErrFormat = errors.New("want dollars.cents, such as 12.34")

// Parse reads "12.34" into 1234 cents. Negative amounts are not allowed.
func Parse(s string) (int, error) {
	whole, frac, ok := strings.Cut(s, ".")
	if !ok || len(frac) != 2 || whole == "" || len(whole) > maxWhole {
		return 0, ErrFormat
	}
	for _, part := range []string{whole, frac} {
		for _, r := range part {
			if r < '0' || r > '9' {
				return 0, ErrFormat
			}
		}
	}
	d, _ := strconv.Atoi(whole) // at most maxWhole digits, all checked above
	c, _ := strconv.Atoi(frac)
	return d*100 + c, nil
}

// Format writes cents as dollars.cents.
func Format(cents int) string {
	return fmt.Sprintf("%d.%02d", cents/100, cents%100)
}
