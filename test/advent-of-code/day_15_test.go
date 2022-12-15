package test

import (
	"advent-of-code-2022/internal/pkg/day15"
	"testing"
)

func TestDay15StarOne(t *testing.T) {
	actual := day15.StarOne("test_input_day_15.txt", 10)

	expected := 26

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

/*
func TestDay15StarTwo(t *testing.T) {
	actual := day14.StarTwo("test_input_day_14.txt")

	expected := 93

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
*/
