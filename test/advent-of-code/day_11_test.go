package test

import (
	"advent-of-code-2022/internal/pkg/day11"
	"testing"
)

func TestDay11StarOne(t *testing.T) {
	actual := day11.StarOne("test_input_day_11.txt")

	expected := int64(10605)

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay11StarTwo(t *testing.T) {
	actual := day11.StarTwo("test_input_day_11.txt")

	expected := int64(2713310158)

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
