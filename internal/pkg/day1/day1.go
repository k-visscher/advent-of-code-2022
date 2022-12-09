package day1

import (
	"advent-of-code-2022/internal/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

func StarOne(input_path string) int {
	input := utils.MustReadFileToString(input_path)

	lines := strings.Split(input, "\n")

	max := 0
	sum := 0
	for _, line := range lines {
		if line == "" {
			if sum >= max {
				max = sum
			}
			sum = 0
			continue
		}

		val, err := strconv.Atoi(line)
		utils.CheckForErr(err)

		sum += val
	}

	return max
}

func StarTwo(input_path string) int {
	input := utils.MustReadFileToString(input_path)

	lines := strings.Split(input, "\n")

	maxes := make([]int, 0)
	sum := 0
	for _, line := range lines {
		if line == "" {
			if len(maxes) < 3 {
				maxes = append(maxes, sum)
				sort.Slice(maxes, func(a, b int) bool {
					return maxes[b] > maxes[a]
				})
			} else {
				for i, max := range maxes {
					if sum >= max {
						maxes[i] = sum
						sort.Slice(maxes, func(a, b int) bool {
							return maxes[b] > maxes[a]
						})
						break
					}
				}
			}

			sum = 0
			continue
		}

		val, err := strconv.Atoi(line)
		utils.CheckForErr(err)

		sum += val
	}

	if sum != 0 {
		for i, max := range maxes {
			if sum >= max {
				maxes[i] = sum
				sort.Slice(maxes, func(a, b int) bool {
					return maxes[b] > maxes[a]
				})
				break
			}
		}
	}

	sum = 0
	for _, max := range maxes {
		sum += max
	}

	return sum
}
