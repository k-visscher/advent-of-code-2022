package test

import (
	"advent-of-code-2022/internal/pkg/day1"
	"testing"
)

func TestDay1StarOne(t *testing.T) {
	actual := day1.StarOne("data/test_input_day_1.txt")

	expected := 24000

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay1StarTwo(t *testing.T) {
	actual := day1.StarTwo("data/test_input_day_1.txt")

	expected := 45000

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
