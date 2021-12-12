package main

import (
	"fmt"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func getAdjacent(m map[int][]int, x int, y int) []int {
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
	for row := 0; row < len(heightMap); row++ {
		for col := range heightMap[row] {
			point := heightMap[row][col]
			adjNums := getAdjacent(heightMap, row, col)
			if isLowPoint(point, adjNums) {
				risks += point + 1
			}
		}
	}
	fmt.Printf("Part 1: %d\n", risks)
}
