package main

import (
	"fmt"
	"sort"

	"github.com/zisk/aoc2021/util"
)

func pop(stack []rune) ([]rune, rune) {
	i := len(stack) - 1
	return stack[:i], stack[i]
}

func checkClose(symbol rune) bool {
	closeList := ")]}>"
	for _, b := range closeList {
		if symbol == b {
			return true
		}
	}
	return false
}

func checkMatch(open rune, close rune) bool {
	switch open {
	case '(':
		if close == ')' {
			return true
		}
	case '[':
		if close == ']' {
			return true
		}
	case '{':
		if close == '}' {
			return true
		}
	case '<':
		if close == '>' {
			return true
		}
	}

	return false
}

func checkCorrupt(s string) (bool, rune) {
	var stack []rune
	var open rune
	for _, sym := range s {
		if !checkClose(sym) {
			stack = append(stack, sym)
			continue
		}
		stack, open = pop(stack)
		if !checkMatch(open, sym) {
			return true, sym
		}
	}
	return false, open
}

func makeComplete(s string) []rune {
	var closeChars []rune
	var stack []rune

	closer := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	for _, sym := range s {
		if !checkClose(sym) {
			stack = append(stack, sym)
			continue
		}
		stack, _ = pop(stack)
	}

	for i := len(stack) - 1; i >= 0; i-- {
		closeChars = append(closeChars, closer[stack[i]])
	}

	return closeChars
}

func main() {
	input, _ := util.InputToTxt()

	corruptScoreMap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	compMap := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	corruptScore := 0
	var incompleteScores []int

	for _, line := range input {
		corrupt, badChar := checkCorrupt(line)
		if corrupt {
			corruptScore += corruptScoreMap[badChar]
		} else {
			closers := makeComplete(line)
			closeScore := 0
			for _, c := range closers {
				closeScore = (closeScore * 5) + compMap[c]
			}
			incompleteScores = append(incompleteScores, closeScore)
		}
	}
	sort.Ints(incompleteScores)

	fmt.Printf("Part 1: %d\n", corruptScore)
	fmt.Printf("Part 2: %d\n", incompleteScores[len(incompleteScores)/2])
}
