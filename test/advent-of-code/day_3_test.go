package test

import (
	"advent-of-code-2022/internal/pkg/day3"
	"testing"
)

func TestDay3StarOne(t *testing.T) {
	actual := day3.StarOne("data/test_input_day_3.txt")

	expected := 157

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay3StarTwo(t *testing.T) {
	actual := day3.StarTwo("data/test_input_day_3.txt")

	expected := 70

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
