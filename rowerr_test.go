package rowerr

import (
	"errors"
	"testing"
)

func TestRowErr(t *testing.T) {
	exp := errors.New("KABOOM")

	if got := New(exp).Scan(); got != exp {
		t.Errorf("expected: %q, got: %q", exp, got)
	}
}
