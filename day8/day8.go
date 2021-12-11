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
	rosetta := make(map[string]string)
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

	ninestr, holding := findNine(holding, numMap[1])
	keyMap[SortString(ninestr)] = "9"
	numMap[9] = ninestr

	rosetta["a"] = findA(numMap[7], numMap[1])
	rosetta["e"] = findE(numMap[8], numMap[9])

	twostr, holding := findTwo(holding, rosetta["e"])
	keyMap[SortString(twostr)] = "2"
	numMap[2] = twostr

	rosetta["g"] = findG(numMap[4], numMap[9], rosetta["a"])
	rosetta["b"] = findB(numMap[8], numMap[2], numMap[1])

	threeStr, fiveStr, holding := findThreeFive(rosetta["b"], holding)
	keyMap[SortString(threeStr)] = "3"
	numMap[3] = threeStr
	keyMap[SortString(fiveStr)] = "5"
	numMap[5] = fiveStr

	rosetta["c"] = findC(numMap[8], numMap[5], rosetta["e"])
	sixStr, zeroStr := findSixZero(rosetta["c"], holding)
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

func findNine(nlist []string, one string) (string, []string) {
	for i, n := range nlist {
		if len(n) == 6 {
			if strings.Contains(n, string(one[0])) && strings.Contains(n, string(one[1])) {
				nlist := append(nlist[:i], nlist[i+1:]...)
				return n, nlist
			}
		}
	}
	return "", nil
}

func findTwo(nlist []string, ePos string) (string, []string) {
	for i, n := range nlist {
		if len(n) == 5 && strings.Contains(n, ePos) {
			nlist := append(nlist[:i], nlist[i+1:]...)
			return n, nlist
		}
	}
	return "", nil
}

func findA(seven string, one string) string {
	s := strings.Replace(seven, string(one[0]), "", -1)
	s = strings.Replace(s, string(one[1]), "", -1)
	return s
}

func findE(eight string, nine string) string {
	for i := range nine {
		eight = strings.Replace(eight, string(nine[i]), "", -1)
	}
	return eight
}

func findG(four string, nine string, aPos string) string {
	n := strings.Replace(nine, aPos, "", -1)
	for i := range four {
		n = strings.Replace(n, string(four[i]), "", -1)
	}
	return n
}

func findB(eight string, two string, one string) string {
	for i := range two {
		eight = strings.Replace(eight, string(two[i]), "", -1)
	}
	for i := range one {
		eight = strings.Replace(eight, string(one[i]), "", -1)
	}
	return eight
}

func findC(eight string, five string, posE string) string {
	eight = strings.Replace(eight, posE, "", -1)
	for i := range five {
		eight = strings.Replace(eight, string(five[i]), "", -1)
	}
	return eight
}

func findThreeFive(posB string, hold []string) (string, string, []string) {
	var three, five string
	for i := len(hold) - 1; i >= 0; i-- {
		if len(hold[i]) == 5 {
			if strings.Contains(hold[i], posB) {
				five = hold[i]
				hold = append(hold[:i], hold[i+1:]...)
			} else {
				three = hold[i]
				hold = append(hold[:i], hold[i+1:]...)
			}
		}
	}
	return three, five, hold
}

func findSixZero(posC string, hold []string) (string, string) {
	var six, zero string
	for i := range hold {
		if strings.Contains(hold[i], posC) {
			six = hold[i]
		} else {
			zero = hold[i]
		}
	}
	return six, zero
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
