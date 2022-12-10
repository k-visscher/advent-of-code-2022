package day10

import (
	"advent-of-code-2022/internal/pkg/utils"
	"fmt"
	"regexp"
	"strings"
)

type InstructionType int

const (
	Noop InstructionType = 1
	Add  InstructionType = 2
)

type Instruction struct {
	Type      InstructionType
	CycleCost int
	TargetVal int
	TargetReg string
}

func StarOne(input_path string) int {
	input := utils.MustReadFileToString(input_path)

	lines := strings.Split(input, "\n")

	regs := make(map[string]int)
	instrs := make([]Instruction, 0)

	measurement_cycles := []int{20, 60, 100, 140, 180, 220}
	measurements := []int{0, 0, 0, 0, 0, 0}

	add_instr_expr := regexp.MustCompile(`add([a-z]) (-?\d+)`)
	nop_instr_expr := regexp.MustCompile(`noop`)

	current_cycle := 0

	for _, line := range lines {
		if add_instr_expr.MatchString(line) {
			expr_groups := add_instr_expr.FindStringSubmatch(line)[1:]

			target := expr_groups[0]
			amount := expr_groups[1]

			instr := Instruction{
				Type:      Add,
				CycleCost: 2,
				TargetVal: utils.MustParseAsInt(amount),
				TargetReg: target,
			}

			regs[instr.TargetReg] = 1

			instrs = append(instrs, instr)
		} else if nop_instr_expr.MatchString(line) {
			instr := Instruction{
				Type:      Noop,
				CycleCost: 1,
			}

			instrs = append(instrs, instr)
		}
	}

	for _, instr := range instrs {
		for cycle := 0; cycle < instr.CycleCost; cycle++ {
			current_cycle++

			current_cycle_is_measurement_cycle := false
			measurement_index := 0
			for i, measurement_cycle := range measurement_cycles {
				measurement_index = i
				if current_cycle == measurement_cycle {
					current_cycle_is_measurement_cycle = true
					break
				}
			}

			if current_cycle_is_measurement_cycle {
				measurements[measurement_index] = current_cycle * regs["x"]
			}

		}

		if instr.Type == Add {
			regs[instr.TargetReg] += instr.TargetVal
		}
	}

	result := 0
	for _, measurement := range measurements {
		result += measurement
	}

	return result
}

func StarTwo(input_path string) string {
	input := utils.MustReadFileToString(input_path)

	lines := strings.Split(input, "\n")

	regs := make(map[string]int)
	instrs := make([]Instruction, 0)

	crt := make([]rune, 240)
	for i := range crt {
		crt[i] = '.'
	}

	add_instr_expr := regexp.MustCompile(`add([a-z]) (-?\d+)`)
	nop_instr_expr := regexp.MustCompile(`noop`)

	current_cycle := 0

	for _, line := range lines {
		if add_instr_expr.MatchString(line) {
			expr_groups := add_instr_expr.FindStringSubmatch(line)[1:]

			target := expr_groups[0]
			amount := expr_groups[1]

			instr := Instruction{
				Type:      Add,
				CycleCost: 2,
				TargetVal: utils.MustParseAsInt(amount),
				TargetReg: target,
			}

			regs[instr.TargetReg] = 1

			instrs = append(instrs, instr)
		} else if nop_instr_expr.MatchString(line) {
			instr := Instruction{
				Type:      Noop,
				CycleCost: 1,
			}

			instrs = append(instrs, instr)
		}
	}

	for _, instr := range instrs {
		for cycle := 0; cycle < instr.CycleCost; cycle++ {
			sprite_pos := regs["x"]

			crt_row_pos := current_cycle % 40
			if crt_row_pos >= sprite_pos-1 && crt_row_pos <= sprite_pos+1 {
				crt[current_cycle] = '#'
			}

			current_cycle++
		}

		if instr.Type == Add {
			regs[instr.TargetReg] += instr.TargetVal
		}
	}

	result := ""
	for i := 0; i < len(crt)-1; i += 40 {
		result += fmt.Sprintf("%s\n", string(crt[i:i+40]))
	}

	return result
}
