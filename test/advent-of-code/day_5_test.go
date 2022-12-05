package test

import (
	"advent-of-code-2022/internal/pkg/day5"
	"testing"
)

func TestDay5StarOne(t *testing.T) {
	actual := day5.StarOne("data/test_input_day_5.txt")

	expected := "CMZ"

	if actual != expected {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}

func TestDay5StarTwo(t *testing.T) {
	actual := day5.StarTwo("data/test_input_day_5.txt")

	expected := "MCD"

	if actual != expected {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}
