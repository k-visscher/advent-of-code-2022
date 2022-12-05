package day5

import (
	"advent-of-code-2022/internal/pkg/utils"
	"regexp"
	"strconv"
	"strings"
)

/*
--- Day 5: Supply Stacks ---

The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2

In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3

In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3

Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3

Finally, one crate is moved from stack 1 to stack 2:

        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3

The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each stack?
*/

func StarOne(input_path string) string {
	input, err := utils.ReadFileToString(input_path)
	utils.CheckForErr(err)

	lines := strings.Split(input, "\n")

	stacks := make(map[int][]string)
	stack_pos_look_up := make(map[int][]int)

	stack_index_expr := regexp.MustCompile(`\d+`)
	procedure_expr := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for line_index, line := range lines {

		if procedure_expr.MatchString(line) {
			groups := procedure_expr.FindStringSubmatch(line)[1:]

			procedureInfo := make([]int, 0)

			for _, group := range groups {
				val, err := strconv.Atoi(group)
				utils.CheckForErr(err)

				procedureInfo = append(procedureInfo, val)
			}

			amount := procedureInfo[0]
			from_stack_index := procedureInfo[1]
			to_stack_index := procedureInfo[2]

			crates_to_move := stacks[from_stack_index][:amount]
			stacks[from_stack_index] = stacks[from_stack_index][amount:]

			for _, crate_to_move := range crates_to_move {
				stacks[to_stack_index] = append([]string{crate_to_move}, stacks[to_stack_index]...)
			}
		} else if stack_index_expr.MatchString(line) {
			matches := stack_index_expr.FindAllStringSubmatchIndex(line, -1)

			for _, match := range matches {
				val, err := strconv.Atoi(line[match[0]:match[1]])
				utils.CheckForErr(err)

				stacks[val] = make([]string, 0)
				stack_pos_look_up[val] = match
			}

			for i := 0; i <= line_index-1; i++ {
				line := lines[i]

				for k, v := range stack_pos_look_up {
					val := line[v[0]:v[1]]
					if val != " " {
						stacks[k] = append(stacks[k], val)
					}
				}
			}
		}
	}

	result := ""
	for i := 1; i <= len(stacks); i++ {
		result += stacks[i][0]
	}

	return result
}

/*
--- Part Two ---

As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.

Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

The CrateMover 9001 is notable for many new and exciting features: air conditioning, leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.

Again considering the example above, the crates begin in the same configuration:

	[D]

[N] [C]
[Z] [M] [P]

	1   2   3

Moving a single crate from stack 2 to stack 1 behaves the same as before:

[D]
[N] [C]
[Z] [M] [P]

	1   2   3

However, the action of moving three crates from stack 1 to stack 3 means that those three moved crates stay in the same order, resulting in this new configuration:

	       [D]
	       [N]
	   [C] [Z]
	   [M] [P]
	1   2   3

Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

	[D]
	[N]

[C]     [Z]
[M]     [P]

	1   2   3

Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

	[D]
	[N]
	[Z]

[M] [C] [P]

	1   2   3

In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.

Before the rearrangement process finishes, update your simulation so that the Elves know where they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?
*/
func StarTwo(input_path string) string {
	input, err := utils.ReadFileToString(input_path)
	utils.CheckForErr(err)

	lines := strings.Split(input, "\n")

	stacks := make(map[int][]string)
	stack_pos_look_up := make(map[int][]int)

	stack_index_expr := regexp.MustCompile(`\d+`)
	procedure_expr := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for line_index, line := range lines {

		if procedure_expr.MatchString(line) {
			groups := procedure_expr.FindStringSubmatch(line)[1:]

			procedureInfo := make([]int, 0)

			for _, group := range groups {
				val, err := strconv.Atoi(group)
				utils.CheckForErr(err)

				procedureInfo = append(procedureInfo, val)
			}

			amount := procedureInfo[0]
			from_stack_index := procedureInfo[1]
			to_stack_index := procedureInfo[2]

			crates_to_move := stacks[from_stack_index][:amount]

			new_to := make([]string, 0)
			new_to = append(new_to, crates_to_move...)
			new_to = append(new_to, stacks[to_stack_index]...)
			stacks[to_stack_index] = new_to

			stacks[from_stack_index] = stacks[from_stack_index][amount:]
		} else if stack_index_expr.MatchString(line) {
			matches := stack_index_expr.FindAllStringSubmatchIndex(line, -1)

			for _, match := range matches {
				val, err := strconv.Atoi(line[match[0]:match[1]])
				utils.CheckForErr(err)

				stacks[val] = make([]string, 0)
				stack_pos_look_up[val] = match
			}

			for i := 0; i <= line_index-1; i++ {
				line := lines[i]

				for k, v := range stack_pos_look_up {
					val := line[v[0]:v[1]]
					if val != " " {
						stacks[k] = append(stacks[k], val)
					}
				}
			}
		}
	}

	result := ""
	for i := 1; i <= len(stacks); i++ {
		result += stacks[i][0]
	}

	return result
}
