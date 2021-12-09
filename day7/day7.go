package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func calcFuel(dist int) int {
	return dist * (dist + 1) / 2
}

func makeRange(low int, high int) []int {
	var result []int
	for i := low; i <= high; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	input, _ := util.InputRaw()
	inSplit := util.StrsToInts(strings.Split(input, ","))
	sort.Ints(inSplit)

	// Part 1
	median := inSplit[len(inSplit)/2]
	fuelCost := 0
	for _, n := range inSplit {
		distToMedian := n - median
		abs := math.Abs(float64(distToMedian))
		fuelCost += int(abs) //int -> float -> int just cause
	}
	fmt.Printf("Part 1\nBest Pos: %d\nFuel: %d\n\n", median, fuelCost)

	// Part 2

	uniquePos := makeRange(inSplit[0], inSplit[len(inSplit)-1])
	var shortPath int
	fuelNeeded := math.MaxInt
	for _, p := range uniquePos {
		var tempFuelCost int
		for _, n := range inSplit {
			distToTry := n - p
			distAbs := math.Abs(float64(distToTry))
			tempFuelCost += calcFuel(int(distAbs))
		}
		if tempFuelCost < fuelNeeded {
			fuelNeeded = tempFuelCost
			shortPath = p
		}
	}

	fmt.Printf("Part 2\nBest Pos: %d\nFuel: %d\n", shortPath, fuelNeeded)
}
