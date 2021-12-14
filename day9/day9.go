package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zisk/aoc2021/util"
)

type point struct {
	x int
	y int
}

func getAdjacentValues(m map[int][]int, x int, y int) []int {
	var r []int
	// up
	if x > 0 {
		r = append(r, m[x-1][y])
	}
	// right
	if (y + 1) < len(m[x]) {
		r = append(r, m[x][y+1])
	}
	// left
	if y > 0 {
		r = append(r, m[x][y-1])
	}
	// down
	if (x + 1) < len(m) {
		r = append(r, m[x+1][y])
	}

	return r
}

func getAdjacentNodes(m map[int][]int, p point) []point {
	var r []point
	// up
	if p.x > 0 {
		r = append(r, point{x: p.x - 1, y: p.y})
	}
	// right
	if (p.y + 1) < len(m[p.x]) {
		r = append(r, point{x: p.x, y: p.y + 1})
	}
	// left
	if p.y > 0 {
		r = append(r, point{x: p.x, y: p.y - 1})
	}
	// down
	if (p.x + 1) < len(m) {
		r = append(r, point{x: p.x + 1, y: p.y})
	}
	return r
}

func popNode(s []point) ([]point, point) {
	l := len(s)
	n := s[l-1]
	if l == 1 {
		return make([]point, 0), n
	}
	return s[:len(s)-1], n
}

func checkVisited(v []point, n point) bool {
	for _, p := range v {
		if p.x == n.x && p.y == n.y {
			return true
		}
	}
	return false
}

func mapBasin(m map[int][]int, start point) []point {
	var next point
	stack := make([]point, 1)
	stack[0] = start
	visited := make([]point, 0)
	for len(stack) > 0 {
		stack, next = popNode(stack)
		newNodes := getAdjacentNodes(m, next)
		for _, node := range newNodes {
			nodeVal := m[node.x][node.y]
			if !(checkVisited(visited, node)) && nodeVal != 9 {
				stack = append(stack, node)
			}
		}
		if !(checkVisited(visited, next)) {
			visited = append(visited, next)
		}
	}
	return visited
}

func getTopBasins(m map[int][]int, basins [][]point) []int {
	var h []int
	for _, p := range basins {
		h = append(h, len(p))
	}
	sort.Ints(h)
	return h[len(h)-3:]
}

func isLowPoint(v int, adj []int) bool {
	for _, i := range adj {
		if v >= i {
			return false
		}
	}
	return true
}

func main() {
	input, _ := util.InputToTxt()
	heightMap := make(map[int][]int)
	for i, r := range input {
		heightMap[i] = util.StrsToInts(strings.Split(r, ""))
	}

	risks := 0
	lowPoints := make([]point, 0)
	for row := 0; row < len(heightMap); row++ {
		for col := range heightMap[row] {
			p := heightMap[row][col]
			adjNums := getAdjacentValues(heightMap, row, col)
			if isLowPoint(p, adjNums) {
				risks += p + 1
				lowPoints = append(lowPoints, point{x: row, y: col})
			}
		}
	}
	fmt.Printf("Part 1: %d\n", risks)

	var basins [][]point
	for _, lowPoint := range lowPoints {
		b := mapBasin(heightMap, lowPoint)
		basins = append(basins, b)
	}
	topThree := getTopBasins(heightMap, basins)
	basinTotal := topThree[0] * topThree[1] * topThree[2]
	fmt.Printf("Top Three Basins: %v\n", topThree)
	fmt.Printf("Part 2: %d\n", basinTotal)
}
