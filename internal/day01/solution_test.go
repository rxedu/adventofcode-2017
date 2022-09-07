package day01

import (
	"testing"
)

func TestSolve(t *testing.T) {
	got := Solve("")
	if got != "" {
		t.Errorf("Solve() = %v; want", got)
	}
}
