package test

import (
	"advent-of-code-2022/internal/pkg/day10"
	"strings"
	"testing"
)

func TestDay10StarOne(t *testing.T) {
	actual := day10.StarOne("test_input_day_10.txt")

	expected := 13140

	if actual != expected {
		t.Fatalf("expected: %d, got: %d", expected, actual)
	}
}

func TestDay10StarTwo(t *testing.T) {
	actual := day10.StarTwo("test_input_day_10.txt")

	expected := `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	expected = strings.TrimSpace(expected)

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}
