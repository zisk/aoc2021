package main

import (
	"fmt"

	"github.com/zisk/aoc2021/util"
)

func main() {
	input, _ := util.InputToInts()
	sonara := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			sonara++
		}
	}

	sonarb := 0
	for i := 3; i < len(input); i++ {
		cur_seq := input[i] + input[i-1] + input[i-2]
		prev_seq := input[i-1] + input[i-2] + input[i-3]
		if cur_seq > prev_seq {
			sonarb++
		}

	}

	fmt.Printf("Part 1: %d\n\n", sonara)
	fmt.Printf("Part 2: %d\n", sonarb)
}
