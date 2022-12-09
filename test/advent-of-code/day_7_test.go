package test

import (
	"advent-of-code-2022/internal/pkg/day7"
	"testing"
)

func TestDay7StarOne(t *testing.T) {
	actual := day7.StarOne("test_input_day_7.txt")

	expected := 95437

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay7StarTwo(t *testing.T) {
	actual := day7.StarTwo("test_input_day_7.txt")

	expected := 24933642

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
