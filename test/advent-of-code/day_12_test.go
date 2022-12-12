package test

import (
	"advent-of-code-2022/internal/pkg/day12"
	"testing"
)

func TestDay12StarOne(t *testing.T) {
	actual := day12.StarOne("test_input_day_12.txt")

	expected := 31

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay12StarTwo(t *testing.T) {
	actual := day12.StarTwo("test_input_day_12.txt")

	expected := 29

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
