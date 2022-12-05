package main

import (
	"advent-of-code-2022/internal/pkg/day1"
	"advent-of-code-2022/internal/pkg/day2"
	"advent-of-code-2022/internal/pkg/day3"
	"advent-of-code-2022/internal/pkg/day4"
	"fmt"
)

func main() {
	fmt.Printf("day 1, star 1: %d\n", day1.StarOne("assets/input_day_1.txt"))
	fmt.Printf("day 1, star 2: %d\n", day1.StarTwo("assets/input_day_1.txt"))

	fmt.Printf("day 2, star 1: %d\n", day2.StarOne("assets/input_day_2.txt"))
	fmt.Printf("day 2, star 1: %d\n", day2.StarTwo("assets/input_day_2.txt"))

	fmt.Printf("day 3, star 1: %d\n", day3.StarOne("assets/input_day_3.txt"))
	fmt.Printf("day 3, star 2: %d\n", day3.StarTwo("assets/input_day_3.txt"))

	fmt.Printf("day 4, star 1: %d\n", day4.StarOne("assets/input_day_4.txt"))
	fmt.Printf("day 4, star 2: %d\n", day4.StarTwo("assets/input_day_4.txt"))
}