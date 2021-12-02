package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func parseCommand(cmd string) (string, int) {
	cmdsplit := strings.Split(cmd, " ")
	dist, _ := strconv.Atoi(cmdsplit[1])
	return cmdsplit[0], dist
}

func main() {

	input, _ := util.InputToTxt()

	// part 1
	horizon := 0
	depth := 0

	for _, command := range input {
		direction, distance := parseCommand(command)

		switch direction {
		case "forward":
			horizon += distance
		case "down":
			depth += distance
		case "up":
			depth -= distance
		}
	}
	fmt.Printf("Horizon: %d\nDepth: %d\n", horizon, depth)
	fmt.Printf("Part 1: %d\n\n", horizon*depth)

	// part 2
	horizon = 0
	depth = 0
	aim := 0

	for _, command := range input {
		direction, distance := parseCommand(command)

		switch direction {
		case "forward":
			horizon += distance
			depth += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		}
	}

	fmt.Printf("Horizon: %d\nDepth: %d\nAim: %d\n", horizon, depth, aim)
	fmt.Printf("Part 2: %d\n", horizon*depth)

}
