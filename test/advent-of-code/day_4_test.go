package test

import (
	"advent-of-code-2022/internal/pkg/day4"
	"testing"
)

func TestDay4StarOne(t *testing.T) {
	actual := day4.StarOne("test_input_day_4.txt")

	expected := 2

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay4StarTwo(t *testing.T) {
	actual := day4.StarTwo("test_input_day_4.txt")

	expected := 4

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
