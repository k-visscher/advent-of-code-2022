package day11

import (
	"advent-of-code-2022/internal/pkg/utils"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type TestOperator string

const (
	Divisible TestOperator = "divisible"
)

type WorryOperator string

const (
	Add      WorryOperator = "+"
	Multiply WorryOperator = "*"
)

type WorryOperandType int

const (
	Variable WorryOperandType = 1
	Constant WorryOperandType = 2
)

type WorryOperand struct {
	Type     WorryOperandType
	Variable string
	Constant int
}

type Monkey struct {
	Items []int64

	TestOperator TestOperator
	TestOperand  int
	TestTargets  map[bool]int

	Inspections int64

	WorryOperator WorryOperator

	WorryOperand WorryOperand
}

func (m *Monkey) PlayRound(monkeys *map[int]*Monkey, modulus int64) {
	for _, oldItem := range m.Items {
		m.Inspections++

		newItem := m.Inspect(oldItem, modulus)
		targetIndex := m.TestTargets[m.Test(newItem)]

		m.Throw(newItem, (*monkeys)[targetIndex])
	}
	m.Items = make([]int64, 0)
}

func (m *Monkey) Inspect(item int64, modulus int64) int64 {
	level := item

	if m.WorryOperand.Type == Constant {
		switch m.WorryOperator {
		case Add:
			level += int64(m.WorryOperand.Constant)
		case Multiply:
			level *= int64(m.WorryOperand.Constant)
		}
	} else {
		switch m.WorryOperator {
		case Add:
			level += level
		case Multiply:
			level *= level
		}
	}

	if modulus == -1 {
		return level / 3
	}

	return level % modulus
}

func (m *Monkey) Test(item int64) bool {
	return item%int64(m.TestOperand) == 0
}

func (*Monkey) Throw(item int64, o *Monkey) {
	o.Items = append(o.Items, item)
}

func StarOne(input_path string) int64 {
	input := utils.MustReadFileToString(input_path)

	monkeys := make(map[int]*Monkey)

	lines := strings.Split(input, "\n")

	monkey_expr := regexp.MustCompile(`Monkey (\d+):`)
	items_expr := regexp.MustCompile(`Starting items: (.*)`)
	worry_expr := regexp.MustCompile(`Operation: new = old (.) (.+)`)

	test_expr := regexp.MustCompile(`Test: ([a-z]+) by (\d+)`)
	target_expr := regexp.MustCompile(`If ((?:true)|(?:false)): throw to monkey (\d+)`)

	for i := 0; i < len(lines); i += 7 {
		number := utils.MustParseAsInt(monkey_expr.FindStringSubmatch(lines[i])[1:][0])

		starting_items := make([]int64, 0)
		raw_items := items_expr.FindStringSubmatch(lines[i+1])[1:][0]
		for _, raw_item := range strings.Split(raw_items, ",") {
			starting_items = append(starting_items, int64(utils.MustParseAsInt(strings.TrimSpace(raw_item))))
		}

		worry_groups := worry_expr.FindStringSubmatch(lines[i+2])[1:]
		worry_operator := WorryOperator(worry_groups[0])
		raw_worry_operand := worry_groups[1]

		worry_operand_struct := WorryOperand{}

		worry_operand_constant, worry_operand_constant_err := strconv.Atoi(raw_worry_operand)
		if worry_operand_constant_err != nil {
			worry_operand_struct.Type = Variable
			worry_operand_struct.Variable = raw_worry_operand
		} else {
			worry_operand_struct.Type = Constant
			worry_operand_struct.Constant = worry_operand_constant
		}

		test_groups := test_expr.FindStringSubmatch(lines[i+3])[1:]
		test_operator := TestOperator(test_groups[0])
		test_operand := utils.MustParseAsInt(test_groups[1])

		targets := make(map[bool]int)
		true_target := utils.MustParseAsInt(target_expr.FindStringSubmatch(lines[i+4])[1:][1])
		targets[true] = true_target
		false_target := utils.MustParseAsInt(target_expr.FindStringSubmatch(lines[i+5])[1:][1])
		targets[false] = false_target

		monkeys[number] = &Monkey{
			Items: starting_items,

			Inspections: 0,

			TestOperator: test_operator,
			TestOperand:  test_operand,
			TestTargets:  targets,

			WorryOperator: worry_operator,
			WorryOperand:  worry_operand_struct,
		}
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeys[j].PlayRound(&monkeys, -1)
		}
	}

	monkeyList := make([]*Monkey, 0)
	for i := 0; i < len(monkeys); i++ {
		monkeyList = append(monkeyList, monkeys[i])
	}

	sort.Slice(monkeyList, func(a, b int) bool {
		return monkeyList[a].Inspections > monkeyList[b].Inspections
	})

	return monkeyList[0].Inspections * monkeyList[1].Inspections
}

func StarTwo(input_path string) int64 {
	input := utils.MustReadFileToString(input_path)

	monkeys := make(map[int]*Monkey)

	lines := strings.Split(input, "\n")

	monkey_expr := regexp.MustCompile(`Monkey (\d+):`)
	items_expr := regexp.MustCompile(`Starting items: (.*)`)
	worry_expr := regexp.MustCompile(`Operation: new = old (.) (.+)`)

	test_expr := regexp.MustCompile(`Test: ([a-z]+) by (\d+)`)
	target_expr := regexp.MustCompile(`If ((?:true)|(?:false)): throw to monkey (\d+)`)

	for i := 0; i < len(lines); i += 7 {
		number := utils.MustParseAsInt(monkey_expr.FindStringSubmatch(lines[i])[1:][0])

		starting_items := make([]int64, 0)
		raw_items := items_expr.FindStringSubmatch(lines[i+1])[1:][0]
		for _, raw_item := range strings.Split(raw_items, ",") {
			starting_items = append(starting_items, int64(utils.MustParseAsInt(strings.TrimSpace(raw_item))))
		}

		worry_groups := worry_expr.FindStringSubmatch(lines[i+2])[1:]
		worry_operator := WorryOperator(worry_groups[0])
		raw_worry_operand := worry_groups[1]

		worry_operand_struct := WorryOperand{}

		worry_operand_constant, worry_operand_constant_err := strconv.Atoi(raw_worry_operand)
		if worry_operand_constant_err != nil {
			worry_operand_struct.Type = Variable
			worry_operand_struct.Variable = raw_worry_operand
		} else {
			worry_operand_struct.Type = Constant
			worry_operand_struct.Constant = worry_operand_constant
		}

		test_groups := test_expr.FindStringSubmatch(lines[i+3])[1:]
		test_operator := TestOperator(test_groups[0])
		test_operand := utils.MustParseAsInt(test_groups[1])

		targets := make(map[bool]int)
		true_target := utils.MustParseAsInt(target_expr.FindStringSubmatch(lines[i+4])[1:][1])
		targets[true] = true_target
		false_target := utils.MustParseAsInt(target_expr.FindStringSubmatch(lines[i+5])[1:][1])
		targets[false] = false_target

		monkeys[number] = &Monkey{
			Items: starting_items,

			Inspections: 0,

			TestOperator: test_operator,
			TestOperand:  test_operand,
			TestTargets:  targets,

			WorryOperator: worry_operator,
			WorryOperand:  worry_operand_struct,
		}
	}

	divisibility_product := 1
	for _, monkey := range monkeys {
		divisibility_product *= monkey.TestOperand
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeys[j].PlayRound(&monkeys, int64(divisibility_product))
		}
	}

	monkeyList := make([]*Monkey, 0)
	for i := 0; i < len(monkeys); i++ {
		monkeyList = append(monkeyList, monkeys[i])
	}

	sort.Slice(monkeyList, func(a, b int) bool {
		return monkeyList[a].Inspections > monkeyList[b].Inspections
	})

	return monkeyList[0].Inspections * monkeyList[1].Inspections
}
