package main

import (
	"advent-of-code-2022/internal/pkg/day1"
	"advent-of-code-2022/internal/pkg/day10"
	"advent-of-code-2022/internal/pkg/day11"
	"advent-of-code-2022/internal/pkg/day12"
	"advent-of-code-2022/internal/pkg/day13"
	"advent-of-code-2022/internal/pkg/day14"
	"advent-of-code-2022/internal/pkg/day15"
	"advent-of-code-2022/internal/pkg/day2"
	"advent-of-code-2022/internal/pkg/day3"
	"advent-of-code-2022/internal/pkg/day4"
	"advent-of-code-2022/internal/pkg/day5"
	"advent-of-code-2022/internal/pkg/day6"
	"advent-of-code-2022/internal/pkg/day7"
	"advent-of-code-2022/internal/pkg/day8"
	"advent-of-code-2022/internal/pkg/day9"
	"fmt"
	"time"
)

func Time[T any](
	to_time func() T,
	after_time func(result T, duration time.Duration)) {
	start := time.Now()
	result := to_time()
	duration := time.Since(start)

	after_time(result, duration)
}

func main() {
	Time(func() int {
		return day1.StarOne("input_day_1.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 1, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day1.StarTwo("input_day_1.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 1, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day2.StarOne("input_day_2.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 2, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day2.StarTwo("input_day_2.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 2, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day3.StarOne("input_day_3.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 3, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day3.StarTwo("input_day_3.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 3, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day4.StarOne("input_day_4.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 4, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day4.StarTwo("input_day_4.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 4, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() string {
		return day5.StarOne("input_day_5.txt")
	}, func(result string, duration time.Duration) {
		fmt.Printf("day 5, star 1: %s, took: %s\n", result, duration)
	})
	Time(func() string {
		return day5.StarTwo("input_day_5.txt")
	}, func(result string, duration time.Duration) {
		fmt.Printf("day 5, star 2: %s, took: %s\n", result, duration)
	})

	Time(func() int {
		return day6.StarOne("input_day_6.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 6, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day6.StarTwo("input_day_6.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 6, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day7.StarOne("input_day_7.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 7, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day7.StarTwo("input_day_7.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 7, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day8.StarOne("input_day_8.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 8, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day8.StarTwo("input_day_8.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 8, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day9.StarOne("input_day_9.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 9, star 1: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day10.StarOne("input_day_10.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 10, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() string {
		return day10.StarTwo("input_day_10.txt")
	}, func(result string, duration time.Duration) {
		fmt.Printf("day 10, star 2:\n%stook: %s\n", result, duration)
	})

	Time(func() int64 {
		return day11.StarOne("input_day_11.txt")
	}, func(result int64, duration time.Duration) {
		fmt.Printf("day 11, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int64 {
		return day11.StarTwo("input_day_11.txt")
	}, func(result int64, duration time.Duration) {
		fmt.Printf("day 11, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day12.StarOne("input_day_12.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 12, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day12.StarTwo("input_day_12.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 12, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day13.StarOne("input_day_13.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 13, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day13.StarTwo("input_day_13.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 13, star 2: %d, took: %s\n", result, duration)
	})

	Time(func() int {
		return day14.StarOne("input_day_14.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 14, star 1: %d, took: %s\n", result, duration)
	})
	Time(func() int {
		return day14.StarTwo("input_day_14.txt")
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 14, star 2: %d, took: %s\n", result, duration)
	})

	// too low = 4033885
	// too high = 5746822
	Time(func() int {
		return day15.StarOne("input_day_15.txt", 2_000_000)
	}, func(result int, duration time.Duration) {
		fmt.Printf("day 15, star 1: %d, took: %s\n", result, duration)
	})
}
