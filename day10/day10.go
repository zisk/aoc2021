package main

import (
	"fmt"

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

func main() {
	input, _ := util.InputToTxt()

	scoreMap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	score := 0

	for _, line := range input {
		corrupt, badChar := checkCorrupt(line)
		if corrupt {
			score += scoreMap[badChar]
		}
	}
	fmt.Printf("Part 1: %d\n", score)
}
