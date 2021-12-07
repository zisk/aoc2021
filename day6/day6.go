package main

import (
	"fmt"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func count(inmap map[int]int) int {
	c := 0
	for _, v := range inmap {
		c += v
	}
	return c
}

func main() {
	input, _ := util.InputRaw()
	inSplit := strings.Split(input, ",")
	inInts := util.StrsToInts(inSplit)

	fishes := make(map[int]int)
	for cycle := 0; cycle <= 8; cycle++ {
		fishes[cycle] = 0
	}

	for _, i := range inInts {
		fishes[i]++
	}

	for i := 0; i < 256; i++ {
		if i == 80 {
			fmt.Printf("Count at 80 Days: %d\n", count(fishes))
		}

		oldFish := make(map[int]int)

		for k, v := range fishes {
			oldFish[k] = v
		}

		for k := 1; k <= 8; k++ {
			fishes[k-1] = fishes[k]
		}
		fishes[6] += oldFish[0]
		fishes[8] = oldFish[0]

	}

	fishCount := count(fishes)

	fmt.Printf("Fish after 256 days: %d\n", fishCount)
}
