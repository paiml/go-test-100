package ledger

import (
	"errors"
	"testing"
)

// fake is a Store held in memory that can be told to fail.
type fake struct {
	xs               []int
	loadErr, saveErr error
}

func (f *fake) Load() ([]int, error) { return f.xs, f.loadErr }
func (f *fake) Save(xs []int) error {
	if f.saveErr != nil {
		return f.saveErr
	}
	f.xs = xs
	return nil
}

func TestAddAndTotal(t *testing.T) {
	f := &fake{}
	l := Ledger{S: f}
	for _, c := range []int{100, 250} {
		if err := l.Add(c); err != nil {
			t.Fatal(err)
		}
	}
	got, err := l.Total()
	if err != nil || got != 350 {
		t.Fatalf("Total = %d, %v", got, err)
	}
}

func TestAddErrors(t *testing.T) {
	boom := errors.New("boom")
	cases := []struct {
		name  string
		f     *fake
		cents int
		want  error
	}{
		{"not positive", &fake{}, 0, nil},
		{"load fails", &fake{loadErr: boom}, 5, boom},
		{"save fails", &fake{saveErr: boom}, 5, boom},
	}
	for _, c := range cases {
		err := Ledger{S: c.f}.Add(c.cents)
		if err == nil {
			t.Errorf("%s: want an error", c.name)
		}
		if c.want != nil && !errors.Is(err, c.want) {
			t.Errorf("%s: err %v does not wrap %v", c.name, err, c.want)
		}
	}
}

func TestTotalLoadError(t *testing.T) {
	boom := errors.New("boom")
	if _, err := (Ledger{S: &fake{loadErr: boom}}).Total(); !errors.Is(err, boom) {
		t.Fatalf("err = %v", err)
	}
}
