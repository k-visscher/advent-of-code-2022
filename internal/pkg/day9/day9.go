package day9

import (
	"advent-of-code-2022/internal/pkg/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Direction string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

func HeadAndTailAreTouching(
	headPosX int, headPosY int,
	tailPosX int, tailPosY int) bool {

	xDelta := int(math.Abs(float64(headPosX - tailPosX)))
	yDelta := int(math.Abs(float64(headPosY - tailPosY)))

	return xDelta <= 1 && yDelta <= 1
}

func HeadAndTailAreInSameRowOrCol(
	headPosX int, headPosY int,
	tailPosX int, tailPosY int) bool {
	return headPosX == tailPosX || headPosY == tailPosY
}

/*
--- Day 9: Rope Bridge ---

This rope bridge creaks as you walk along it. You aren't sure how old it is, or whether it can even support your weight.

It seems to support the Elves just fine, though. The bridge spans a gorge which was carved out by the massive river far below you.

You step carefully; as you do, the ropes stretch and twist. You decide to distract yourself by modeling rope physics; maybe you can even figure out where not to step.

Consider a rope with a knot at each end; these knots mark the head and the tail of the rope. If the head moves far enough away from the tail, the tail is pulled toward the head.

Due to nebulous reasoning involving Planck lengths, you should be able to model the positions of the knots on a two-dimensional grid. Then, by following a hypothetical series of motions (your puzzle input) for the head, you can determine how the tail will move.

Due to the aforementioned Planck lengths, the rope must be quite short; in fact, the head (H) and tail (T) must always be touching (diagonally adjacent and even overlapping both count as touching):

....
.TH.
....

....
.H..
..T.
....

...
.H. (H covers T)
...

If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step in that direction so it remains close enough:

.....    .....    .....
.TH.. -> .T.H. -> ..TH.
.....    .....    .....

...    ...    ...
.T.    .T.    ...
.H. -> ... -> .T.
...    .H.    .H.
...    ...    ...

Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves one step diagonally to keep up:

.....    .....    .....
.....    ..H..    ..H..
..H.. -> ..... -> ..T..
.T...    .T...    .....
.....    .....    .....

.....    .....    .....
.....    .....    .....
..H.. -> ...H. -> ..TH.
.T...    .T...    .....
.....    .....    .....

You just need to work out where the tail goes as the head follows a series of motions. Assume the head and the tail both start at the same position, overlapping.

For example:

R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2

This series of motions moves the head right four steps, then up four steps, then left three steps, then down one step, and so on. After each step, you'll need to update the position of the tail if the step means the head is no longer adjacent to the tail. Visually, these motions occur as follows (s marks the starting position as a reference point):

== Initial State ==

......
......
......
......
H.....  (H covers T, s)

== R 4 ==

......
......
......
......
TH....  (T covers s)

......
......
......
......
sTH...

......
......
......
......
s.TH..

......
......
......
......
s..TH.

== U 4 ==

......
......
......
....H.
s..T..

......
......
....H.
....T.
s.....

......
....H.
....T.
......
s.....

....H.
....T.
......
......
s.....

== L 3 ==

...H..
....T.
......
......
s.....

..HT..
......
......
......
s.....

.HT...
......
......
......
s.....

== D 1 ==

..T...
.H....
......
......
s.....

== R 4 ==

..T...
..H...
......
......
s.....

..T...
...H..
......
......
s.....

......
...TH.
......
......
s.....

......
....TH
......
......
s.....

== D 1 ==

......
....T.
.....H
......
s.....

== L 5 ==

......
....T.
....H.
......
s.....

......
....T.
...H..
......
s.....

......
......
..HT..
......
s.....

......
......
.HT...
......
s.....

......
......
HT....
......
s.....

== R 2 ==

......
......
.H....  (H covers T)
......
s.....

......
......
.TH...
......
s.....

After simulating the rope, you can count up all of the positions the tail visited at least once. In this diagram, s again marks the starting position (which the tail also visited) and # marks other positions the tail visited:

..##..
...##.
.####.
....#.
s###..

So, there are 13 positions the tail visited at least once.

Simulate your complete hypothetical series of motions. How many positions does the tail of the rope visit at least once?
*/
func StarOne(input_path string) int {
	input := utils.MustReadFileToString(input_path)

	lines := strings.Split(input, "\n")

	prevHeadPosX := 0
	prevHeadPosY := 0

	headPosX := 0
	headPosY := 0

	tailPosX := 0
	tailPosY := 0

	move_expr := regexp.MustCompile(`([U|D|L|R]) (\d+)`)
	tail_pos_set := make(map[string]bool) //string as key since we can't do a []int as key
	tail_pos_set[fmt.Sprintf("%d,%d", tailPosX, tailPosY)] = true

	for _, line := range lines {
		move_expr_groups := move_expr.FindStringSubmatch(line)[1:]

		direction := Direction(move_expr_groups[0])
		distance := utils.MustParseAsInt(move_expr_groups[1])

		// coordinate system origin @ bottom left
		switch direction {
		case Up:
			for i := 0; i < distance; i++ {
				prevHeadPosX = headPosX
				prevHeadPosY = headPosY

				headPosY++
				if !HeadAndTailAreTouching(
					headPosX, headPosY,
					tailPosX, tailPosY) {
					if !HeadAndTailAreInSameRowOrCol(
						headPosX, headPosY,
						tailPosX, tailPosY) {
						tailPosX = prevHeadPosX
						tailPosY = prevHeadPosY
					} else {
						tailPosY++
					}

					tail_pos_set[fmt.Sprintf("%d,%d", tailPosX, tailPosY)] = true
				}
			}
		case Left:
			for i := 0; i < distance; i++ {
				prevHeadPosX = headPosX
				prevHeadPosY = headPosY

				headPosX--
				if !HeadAndTailAreTouching(
					headPosX, headPosY,
					tailPosX, tailPosY) {
					if !HeadAndTailAreInSameRowOrCol(
						headPosX, headPosY,
						tailPosX, tailPosY) {
						tailPosX = prevHeadPosX
						tailPosY = prevHeadPosY
					} else {
						tailPosX--
					}

					tail_pos_set[fmt.Sprintf("%d,%d", tailPosX, tailPosY)] = true
				}
			}
		case Right:
			for i := 0; i < distance; i++ {
				prevHeadPosX = headPosX
				prevHeadPosY = headPosY

				headPosX++
				if !HeadAndTailAreTouching(
					headPosX, headPosY,
					tailPosX, tailPosY) {
					if !HeadAndTailAreInSameRowOrCol(
						headPosX, headPosY,
						tailPosX, tailPosY) {
						tailPosX = prevHeadPosX
						tailPosY = prevHeadPosY
					} else {
						tailPosX++
					}

					tail_pos_set[fmt.Sprintf("%d,%d", tailPosX, tailPosY)] = true
				}
			}
		case Down:
			for i := 0; i < distance; i++ {
				prevHeadPosX = headPosX
				prevHeadPosY = headPosY

				headPosY--
				if !HeadAndTailAreTouching(
					headPosX, headPosY,
					tailPosX, tailPosY) {
					if !HeadAndTailAreInSameRowOrCol(
						headPosX, headPosY,
						tailPosX, tailPosY) {
						tailPosX = prevHeadPosX
						tailPosY = prevHeadPosY
					} else {
						tailPosY--
					}

					tail_pos_set[fmt.Sprintf("%d,%d", tailPosX, tailPosY)] = true
				}
			}
		}
	}

	return len(tail_pos_set)
}

/*
func StarTwo(input_path string) int {
	input, err := utils.ReadFileToString(input_path)
	utils.CheckForErr(err)

	lines := strings.Split(input, "\n")

	return 0
}
*/
