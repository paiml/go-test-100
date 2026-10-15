// Package ledger keeps a running total in a Store. Store is an interface so
// tests can swap in a fake that fails on demand.
package ledger

import "fmt"

// Store persists a list of amounts in cents.
type Store interface {
	Load() ([]int, error)
	Save([]int) error
}

// Ledger adds amounts and reports their total.
type Ledger struct{ S Store }

// Add appends cents to the stored list.
func (l Ledger) Add(cents int) error {
	if cents <= 0 {
		return fmt.Errorf("ledger: amount must be positive, got %d", cents)
	}
	xs, err := l.S.Load()
	if err != nil {
		return fmt.Errorf("ledger: load: %w", err)
	}
	if err := l.S.Save(append(xs, cents)); err != nil {
		return fmt.Errorf("ledger: save: %w", err)
	}
	return nil
}

// Total sums the stored amounts.
func (l Ledger) Total() (int, error) {
	xs, err := l.S.Load()
	if err != nil {
		return 0, fmt.Errorf("ledger: load: %w", err)
	}
	t := 0
	for _, x := range xs {
		t += x
	}
	return t, nil
}
