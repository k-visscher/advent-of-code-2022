package test

import (
	"advent-of-code-2022/internal/pkg/day13"
	"testing"
)

func TestDay13StarOne(t *testing.T) {
	actual := day13.StarOne("test_input_day_13.txt")

	expected := 13

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay13StarTwo(t *testing.T) {
	actual := day13.StarTwo("test_input_day_13.txt")

	expected := 140

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
