package main

import (
	"fmt"
	"math"
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

func neg(a int) int {
	if a > 0 {
		return 1
	}
	return -1
}

type point struct {
	x, y int
}

type line struct {
	start point
	end   point
}

func (l *line) trace(diag bool) []point {
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
	} else if diag {
		slopeX := neg(l.end.x - l.start.x)
		slopeY := neg(l.end.y - l.start.y)
		diffX := int(math.Abs(float64(l.start.x) - float64(l.end.x)))

		for p := 0; p <= diffX; p++ {
			distX := l.start.x + (p * slopeX)
			distY := l.start.y + (p * slopeY)
			lineTrace = append(lineTrace, point{x: distX, y: distY})
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

	// Part 1
	var points []point
	for _, ls := range lines {
		newPoints := ls.trace(false)
		if len(newPoints) != 0 {
			points = append(points, newPoints...)
		}
	}

	dupeMap1 := make(map[string]int)

	for _, ps := range points {
		k := fmt.Sprintf("%03d:%03d", ps.x, ps.y)
		dupeMap1[k]++
	}

	overCount1 := 0
	for _, m := range dupeMap1 {
		if m > 1 {
			overCount1++
		}
	}

	// Part 2
	var diagPoints []point
	diagMap := make(map[string]int)
	overCount2 := 0
	for _, ls := range lines {
		ps := ls.trace(true)
		diagPoints = append(diagPoints, ps...)
	}
	for _, ps := range diagPoints {
		k := fmt.Sprintf("%03d:%03d", ps.x, ps.y)
		diagMap[k]++
	}
	for _, m := range diagMap {
		if m > 1 {
			overCount2++
		}
	}

	fmt.Printf("Part 1: %d\n", overCount1)
	fmt.Printf("Part 2: %d\n", overCount2)
}
