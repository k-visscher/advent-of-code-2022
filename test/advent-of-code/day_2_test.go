package test

import (
	"advent-of-code-2022/internal/pkg/day2"
	"testing"
)

func TestDay2StarOne(t *testing.T) {
	actual := day2.StarOne("data/test_input_day_2.txt")

	expected := 15

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay2StarTwo(t *testing.T) {
	actual := day2.StarTwo("data/test_input_day_2.txt")

	expected := 12

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
