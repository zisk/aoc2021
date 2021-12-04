package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func BitPopularity(input []string) map[int][]int {
	bitMap := make(map[int][]int)

	for i := 0; i < len(input[0]); i++ {
		bitMap[i] = []int{0, 0}
	}

	for _, line := range input {
		for i, bit := range strings.Split(line, "") {
			valueToInt, _ := strconv.Atoi(bit)
			bitMap[i][valueToInt]++
		}
	}
	return bitMap
}
func main() {
	input, _ := util.InputToTxt()

	bitCount := BitPopularity(input)

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

	fmt.Printf("\nGamma: %s\nEpsilon: %s\n\n", gamma, epsilon)
	fmt.Printf("Gamma: %d\nEpsilon: %d\n\n", gammaInt, epsilonInt)
	fmt.Printf("Power Consumption: %d\n\n", gammaInt*epsilonInt)

	oxResult := make([]string, len(input))
	_ = copy(oxResult, input)

	for i := 0; i < len(input[0]); i++ {
		oxPop := BitPopularity(oxResult)
		b := 0
		if oxPop[i][0] == oxPop[i][1] {
			b = 1
		} else if oxPop[i][0] < oxPop[i][1] {
			b = 1
		}

		var oxHold []string
		for _, n := range oxResult {
			if strings.Split(n, "")[i] == strconv.Itoa(b) {
				oxHold = append(oxHold, n)
			}
		}

		oxResult = oxHold

		if len(oxResult) == 1 {
			break
		}
	}

	coResult := make([]string, len(input))
	_ = copy(coResult, input)

	for i := 0; i < len(input[0]); i++ {
		coPop := BitPopularity(coResult)
		b := 1
		if coPop[i][0] == coPop[i][1] {
			b = 0
		} else if coPop[i][0] < coPop[i][1] {
			b = 0
		}

		var coHold []string
		for _, n := range coResult {
			if strings.Split(n, "")[i] == strconv.Itoa(b) {
				coHold = append(coHold, n)
			}
		}

		coResult = coHold

		if len(coResult) == 1 {
			break
		}
	}

	oxInt, _ := strconv.ParseInt(oxResult[0], 2, 64)
	coInt, _ := strconv.ParseInt(coResult[0], 2, 64)
	fmt.Printf("Oxygen Rate: %d\n", oxInt)
	fmt.Printf("C02 Rate: %d\n", coInt)
	fmt.Printf("Life support rating: %d\n", oxInt*coInt)
}
