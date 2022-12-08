package test

import (
	"advent-of-code-2022/internal/pkg/day8"
	"testing"
)

func TestDay8StarOne(t *testing.T) {
	actual := day8.StarOne("data/test_input_day_8.txt")

	expected := 21

	t.Logf("%v", actual)

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

/*
func TestDay7StarTwo(t *testing.T) {
	actual := day7.StarTwo("data/test_input_day_7.txt")

	expected := 24933642

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}
*/
