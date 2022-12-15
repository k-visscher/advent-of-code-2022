package day15

import (
	"advent-of-code-2022/internal/pkg/utils"
	"regexp"
	"strings"
)

type SensorBeaconPair struct {
	Sensor utils.Position
	Beacon utils.Position
}

func ManhattanDistance(p1 *utils.Position, p2 *utils.Position) int {
	return utils.Abs(p2.X-p1.X) + utils.Abs(p2.Y-p1.Y)
}

func (sbp *SensorBeaconPair) ToCircle() Circle {
	return Circle{Center: sbp.Sensor, Radius: ManhattanDistance(&sbp.Sensor, &sbp.Beacon)}
}

type Circle struct {
	Center utils.Position
	Radius int
}

// TODO: USE ME, https://en.wikipedia.org/wiki/Taxicab_geometry see section on balls
// https://math.stackexchange.com/a/198767
func (c *Circle) ContainsPosition(p *utils.Position) bool {
	return ManhattanDistance(p, &c.Center) <= c.Radius
}

func StarOne(input_path string, y int) int {
	in := utils.MustReadFileToString(input_path)
	ls := strings.Split(in, "\n")

	sb_xpr := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	sb_prs := make([]SensorBeaconPair, 0)
	cs := make([]Circle, 0)

	min := utils.Position{X: 0, Y: 0}
	max := utils.Position{X: 0, Y: 0}
	first := true

	for _, l := range ls {
		rps := sb_xpr.FindStringSubmatch(l)[1:]

		ps := make([]int, 0)
		for _, rp := range rps {
			ps = append(ps, utils.MustParseAsInt(rp))
		}

		s := utils.Position{X: ps[0], Y: ps[1]}
		b := utils.Position{X: ps[2], Y: ps[3]}
		if first {
			min = s
			max = s
			first = false
		}

		min.X = utils.Min(min.X, s.X)
		max.X = utils.Max(max.X, s.X)
		min.X = utils.Min(min.X, b.X)
		max.X = utils.Max(max.X, b.X)

		sb_pr := SensorBeaconPair{Sensor: s, Beacon: b}
		c := sb_pr.ToCircle()

		sb_prs = append(sb_prs, sb_pr)
		cs = append(cs, c)
	}

	for _, c := range cs {
		min.X = utils.Min(min.X, c.Center.X-c.Radius)
		max.X = utils.Max(max.X, c.Center.X+c.Radius)
	}

	result := 0
	for x := min.X; x <= max.X; x++ {
		cp := utils.Position{X: x, Y: y}

		found_sb_at_cp := false
		for _, sb := range sb_prs {
			if sb.Beacon.X == cp.X && sb.Beacon.Y == cp.Y ||
				sb.Sensor.X == cp.X && sb.Sensor.Y == cp.Y {
				found_sb_at_cp = true
				break
			}
		}
		if found_sb_at_cp {
			continue
		}

		for _, c := range cs {
			if c.ContainsPosition(&cp) {
				result++
				break
			}
		}
	}

	return result
}
