package money

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		in   string
		want int
		err  error
	}{
		{"12.34", 1234, nil}, {"0.05", 5, nil}, {"7.00", 700, nil},
		{"12", 0, ErrFormat}, {"12.3", 0, ErrFormat}, {".50", 0, ErrFormat},
		{"1x.50", 0, ErrFormat}, {"1.5x", 0, ErrFormat}, {"-1.00", 0, ErrFormat},
		{"99999999999999999999.00", 0, ErrFormat},
	}
	for _, c := range cases {
		got, err := Parse(c.in)
		if got != c.want || !errors.Is(err, c.err) {
			t.Errorf("Parse(%q) = %d, %v; want %d, %v", c.in, got, err, c.want, c.err)
		}
	}
}

func TestFormat(t *testing.T) {
	if got := Format(1234); got != "12.34" {
		t.Fatalf("got %q", got)
	}
	if got := Format(5); got != "0.05" {
		t.Fatalf("got %q", got)
	}
}

// FuzzRoundTrip: whatever Parse accepts, Format must give back unchanged.
func FuzzRoundTrip(f *testing.F) {
	for _, s := range []string{"12.34", "0.00", "x", "1.5", "007.01"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		c, err := Parse(s)
		if err != nil {
			return
		}
		back, err := Parse(Format(c))
		if err != nil || back != c {
			t.Fatalf("%q -> %d -> %q -> %d, %v", s, c, Format(c), back, err)
		}
	})
}
