package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func main() {
	input, _ := util.InputToTxt()

	bitCount := make(map[int][]int)

	// populate the lists
	for i := 0; i < len(input[0]); i++ {
		bitCount[i] = []int{0, 0}
	}

	for _, line := range input {
		for i, bit := range strings.Split(line, "") {
			valueToInt, _ := strconv.Atoi(bit)
			bitCount[i][valueToInt]++
		}
	}

	var gamma string
	var epsilon string

	for i := 0; i < len(input[0]); i++ {
		if bitCount[i][0] > bitCount[i][1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(bitCount)
	fmt.Printf("\nGamma: %s\nEpsilon: %s\n\n", gamma, epsilon)
	fmt.Printf("Gamma: %d\nEpsilon: %d\n\n", gammaInt, epsilonInt)
	fmt.Printf("Power Consumption: %d\n", gammaInt*epsilonInt)
}
