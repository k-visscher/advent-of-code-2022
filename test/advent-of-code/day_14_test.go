package test

import (
	"advent-of-code-2022/internal/pkg/day14"
	"testing"
)

func TestDay14StarOne(t *testing.T) {
	actual := day14.StarOne("test_input_day_14.txt")

	expected := 24

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
