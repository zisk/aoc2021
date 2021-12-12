package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

type display struct {
	pattern []string
	output  []string
}

func (d *display) countEasyNums() int {
	count := 0
	for _, n := range d.output {
		switch len(n) {
		case 2:
			count++
		case 3:
			count++
		case 7:
			count++
		case 4:
			count++
		}

	}
	return count
}

func (d *display) decode() int {
	keyMap := make(map[string]string)
	numMap := make(map[int]string)
	holding := make([]string, len(d.pattern))
	copy(holding, d.pattern)

	// first pass
	// find the easy numbers and pop
	for p := len(holding) - 1; p >= 0; p-- {
		switch len(holding[p]) {
		case 7:
			keyMap[SortString(holding[p])] = "8"
			numMap[8] = holding[p]
			holding = append(holding[:p], holding[p+1:]...)
		case 4:
			keyMap[SortString(holding[p])] = "4"
			numMap[4] = holding[p]
			holding = append(holding[:p], holding[p+1:]...)
		case 3:
			keyMap[SortString(holding[p])] = "7"
			numMap[7] = holding[p]
			holding = append(holding[:p], holding[p+1:]...)
		case 2:
			keyMap[SortString(holding[p])] = "1"
			numMap[1] = holding[p]
			holding = append(holding[:p], holding[p+1:]...)
		}
	}

	ninestr, holding := findNine(holding, numMap[4])
	keyMap[SortString(ninestr)] = "9"
	numMap[9] = ninestr

	threeStr, holding := findThree(holding, numMap[7])
	keyMap[SortString(threeStr)] = "3"
	numMap[3] = threeStr

	twostr, fiveStr, holding := findTwoFive(holding, numMap[4])
	keyMap[SortString(twostr)] = "2"
	numMap[2] = twostr
	keyMap[SortString(fiveStr)] = "5"
	numMap[5] = fiveStr

	sixStr, zeroStr := findSixZero(holding, numMap[7])
	keyMap[SortString(sixStr)] = "6"
	numMap[6] = sixStr
	keyMap[SortString(zeroStr)] = "0"
	numMap[0] = zeroStr

	var o string
	for _, n := range d.output {
		r := keyMap[SortString(n)]
		o += r
	}

	decoded, _ := strconv.Atoi(o)
	return decoded

}

func findNine(hold []string, four string) (string, []string) {
	for i := len(hold) - 1; i >= 0; i-- {
		if len(hold[i]) == 6 {
			if checkContains(hold[i], four) {
				nine := hold[i]
				hold = append(hold[:i], hold[i+1:]...)
				return nine, hold
			}
		}
	}
	return "", nil
}

func findThree(hold []string, seven string) (string, []string) {
	for i := len(hold) - 1; i >= 0; i-- {
		if len(hold[i]) == 5 {
			if checkContains(hold[i], seven) {
				three := hold[i]
				hold = append(hold[:i], hold[i+1:]...)
				return three, hold
			}
		}
	}
	return "", nil
}

func findTwoFive(hold []string, four string) (string, string, []string) {
	var five, two string
	for i := len(hold) - 1; i >= 0; i-- {
		if len(hold[i]) == 5 {
			intersectCount := 0
			for _, l := range four {
				s := string(l)
				if strings.Contains(hold[i], s) {
					intersectCount++
				}
			}
			if intersectCount == 3 {
				five = hold[i]
				hold = append(hold[:i], hold[i+1:]...)
			} else {
				two = hold[i]
				hold = append(hold[:i], hold[i+1:]...)
			}
		}
	}
	return two, five, hold
}

func findSixZero(hold []string, seven string) (string, string) {
	var six, zero string
	for i := len(hold) - 1; i >= 0; i-- {
		if len(hold[i]) == 6 {
			if checkContains(hold[i], seven) {
				zero = hold[i]
				hold = append(hold[:i], hold[i+1:]...)
			} else {
				six = hold[i]
				hold = append(hold[:i], hold[i+1:]...)
			}
		}
	}

	return six, zero
}

func checkContains(a string, b string) bool {
	for _, i := range b {
		if !(strings.Contains(a, string(i))) {
			return false
		}
	}
	return true
}

func makeDisplay(in string) display {
	initialSplit := strings.Split(in, " | ")
	return display{pattern: strings.Fields(initialSplit[0]), output: strings.Fields(initialSplit[1])}
}

func main() {
	input, _ := util.InputToTxt()
	var displays []display
	for _, i := range input {
		s := makeDisplay(i)
		displays = append(displays, s)
	}

	// Part 1
	easyNums := 0
	for _, dist := range displays {
		easyNums += dist.countEasyNums()
	}

	fmt.Printf("Part 1: %d\n\n", easyNums)

	// Part 2
	decodedResult := 0
	for _, dist := range displays {
		decodedResult += dist.decode()
	}

	fmt.Printf("Part 2: %d\n", decodedResult)
}
