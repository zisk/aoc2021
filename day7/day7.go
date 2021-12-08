package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func main() {
	input, _ := util.InputRaw()
	inSplit := util.StrsToInts(strings.Split(input, ","))
	sort.Ints(inSplit)
	median := inSplit[len(inSplit)/2]
	fuelCost := 0
	for _, n := range inSplit {
		distToMedian := n - median
		abs := math.Abs(float64(distToMedian))
		fuelCost += int(abs)
	}
	fmt.Printf("Fuel Cost Part 1: %d", fuelCost)

}
