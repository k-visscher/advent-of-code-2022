package test

import (
	"advent-of-code-2022/internal/pkg/day6"
	"testing"
)

func TestDay6StarOne(t *testing.T) {
	actual := day6.StarOne("test_input_day_6.txt")

	expected := 10

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay6StarTwo(t *testing.T) {
	actual := day6.StarTwo("test_input_day_6.txt")

	expected := 29

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
