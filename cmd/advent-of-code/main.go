package main

import (
	"advent-of-code-2022/internal/pkg/day1"
	"advent-of-code-2022/internal/pkg/day10"
	"advent-of-code-2022/internal/pkg/day11"
	"advent-of-code-2022/internal/pkg/day2"
	"advent-of-code-2022/internal/pkg/day3"
	"advent-of-code-2022/internal/pkg/day4"
	"advent-of-code-2022/internal/pkg/day5"
	"advent-of-code-2022/internal/pkg/day6"
	"advent-of-code-2022/internal/pkg/day7"
	"advent-of-code-2022/internal/pkg/day8"
	"advent-of-code-2022/internal/pkg/day9"
	"fmt"
)

func main() {
	fmt.Printf("day 1, star 1: %d\n", day1.StarOne("input_day_1.txt"))
	fmt.Printf("day 1, star 2: %d\n", day1.StarTwo("input_day_1.txt"))

	fmt.Printf("day 2, star 1: %d\n", day2.StarOne("input_day_2.txt"))
	fmt.Printf("day 2, star 2: %d\n", day2.StarTwo("input_day_2.txt"))

	fmt.Printf("day 3, star 1: %d\n", day3.StarOne("input_day_3.txt"))
	fmt.Printf("day 3, star 2: %d\n", day3.StarTwo("input_day_3.txt"))

	fmt.Printf("day 4, star 1: %d\n", day4.StarOne("input_day_4.txt"))
	fmt.Printf("day 4, star 2: %d\n", day4.StarTwo("input_day_4.txt"))

	fmt.Printf("day 5, star 1: %s\n", day5.StarOne("input_day_5.txt"))
	fmt.Printf("day 5, star 2: %s\n", day5.StarTwo("input_day_5.txt"))

	fmt.Printf("day 6, star 1: %d\n", day6.StarOne("input_day_6.txt"))
	fmt.Printf("day 6, star 2: %d\n", day6.StarTwo("input_day_6.txt"))

	fmt.Printf("day 7, star 1: %d\n", day7.StarOne("input_day_7.txt"))
	fmt.Printf("day 7, star 2: %d\n", day7.StarTwo("input_day_7.txt"))

	fmt.Printf("day 8, star 1: %d\n", day8.StarOne("input_day_8.txt"))
	fmt.Printf("day 8, star 2: %d\n", day8.StarTwo("input_day_8.txt"))

	fmt.Printf("day 9, star 1: %d\n", day9.StarOne("input_day_9.txt"))
	//fmt.Printf("day 9, star 2: %d\n", day9.StarTwo("test_input_day_9_part_1.txt"))

	fmt.Printf("day 10, star 1: %d\n", day10.StarOne("input_day_10.txt"))
	fmt.Printf("day 10, star 2: \n%s\n", day10.StarTwo("input_day_10.txt"))

	fmt.Printf("day 11, star 1: %d\n", day11.StarOne("input_day_11.txt"))
	fmt.Printf("day 11, star 2: %d\n", day11.StarTwo("test_input_day_11.txt", nil))
}
