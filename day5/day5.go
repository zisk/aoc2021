package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

type point struct {
	x, y int
}

type line struct {
	start point
	end   point
}

func (l *line) trace() []point {
	var lineTrace []point
	if l.start.x != l.end.x && l.start.y == l.end.y {
		highPoint := max(l.start.x, l.end.x)
		lowPoint := min(l.start.x, l.end.x)
		for p := lowPoint; p <= highPoint; p++ {
			lineTrace = append(lineTrace, point{x: p, y: l.start.y})
		}
	} else if l.start.y != l.end.y && l.start.x == l.end.x {
		highPoint := max(l.start.y, l.end.y)
		lowPoint := min(l.start.y, l.end.y)
		for p := lowPoint; p <= highPoint; p++ {
			lineTrace = append(lineTrace, point{x: l.start.x, y: p})
		}
	}
	return lineTrace
}

func newLine(lineString string) line {
	s := strings.Split(lineString, " ")
	p1 := strings.Split(s[0], ",")
	p2 := strings.Split(s[2], ",")
	p1a, _ := strconv.Atoi(p1[0])
	p1b, _ := strconv.Atoi(p1[1])
	p2a, _ := strconv.Atoi(p2[0])
	p2b, _ := strconv.Atoi(p2[1])
	return line{start: point{p1a, p1b}, end: point{p2a, p2b}}
}

func main() {
	input, _ := util.InputToTxt()
	var lines []line
	for _, line := range input {
		lines = append(lines, newLine(line))
	}

	// fmt.Println(lines)
	var points []point
	for _, ls := range lines {
		newPoints := ls.trace()
		if len(newPoints) != 0 {
			points = append(points, newPoints...)
		}
	}

	dupeMap := make(map[string]int)

	for _, ps := range points {
		k := fmt.Sprintf("%03d:%03d", ps.x, ps.y)
		dupeMap[k]++
	}

	overCount := 0
	for _, m := range dupeMap {
		if m > 1 {
			overCount++
		}
	}

	fmt.Printf("Overlapping Points: %d\n", overCount)
}
