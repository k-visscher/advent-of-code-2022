package test

import (
	"advent-of-code-2022/internal/pkg/day9"
	"testing"
)

func TestDay9StarOne(t *testing.T) {
	actual := day9.StarOne("test_input_day_9_part_1.txt")

	expected := 13

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay9StarTwo(t *testing.T) {
	actual := day9.StarTwo("test_input_day_9_part_2.txt")

	expected := 36

	t.Logf("%v", actual)

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
