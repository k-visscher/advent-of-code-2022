package test

import (
	"advent-of-code-2022/internal/pkg/day8"
	"testing"
)

func TestDay8StarOne(t *testing.T) {
	actual := day8.StarOne("test_input_day_8.txt")

	expected := 21

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay8StarTwo(t *testing.T) {
	actual := day8.StarTwo("test_input_day_8.txt")

	expected := 8

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
