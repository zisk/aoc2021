package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/zisk/aoc2021/util"
)

func parseDraw(draws string) []int {
	drawSplit := strings.Split(draws, ",")
	result := make([]int, len(drawSplit))
	for i, d := range drawSplit {
		s, _ := strconv.Atoi(d)
		result[i] = s
	}
	return result
}

type card struct {
	nums map[int][5]int
	hits map[int]*[5]bool
}

func (c *card) check(draw int) {
search:
	for ro := 0; ro < 5; ro++ {
		for col := 0; col < 5; col++ {
			if c.nums[ro][col] == draw {
				c.hits[ro][col] = true
				break search
			}
		}
	}
}

func (c *card) winner() bool {

	for i := 0; i < 5; i++ {
		rowCount := 0
		for _, col := range c.hits[i] {
			if col {
				rowCount++
			}
			if rowCount == 5 {
				return true
			}
		}
		colCount := 0
		for r := 0; r < 5; r++ {
			if c.hits[r][i] {
				colCount++
			}
			if colCount == 5 {
				return true
			}
		}
	}

	return false
}

func (c *card) score(lastdraw int) int {
	score := 0

	for ro := 0; ro < 5; ro++ {
		for col := 0; col < 5; col++ {
			if !(c.hits[ro][col]) {
				score += c.nums[ro][col]
			}
		}
	}

	return score * lastdraw
}

func makeCard(rawcard string) card {
	rowSplitReg := regexp.MustCompile(`\s+`)
	numGrid := make(map[int][5]int)
	rows := strings.Split(rawcard, "\n")
	for i, row := range rows {
		rowSplit := rowSplitReg.Split(strings.Trim(row, " "), -1)
		var rowInts [5]int
		for p, n := range rowSplit {
			r, _ := strconv.Atoi(n)
			rowInts[p] = r
		}
		numGrid[i] = rowInts
	}
	hits := make(map[int]*[5]bool)
	for i := 0; i < 5; i++ {
		hits[i] = &[5]bool{}
	}
	return card{nums: numGrid, hits: hits}
}

func main() {

	input, _ := util.InputRaw()

	blankLines := regexp.MustCompile("\n\n")
	inSplit := blankLines.Split(input, -1)
	draws := parseDraw(inSplit[0])
	rawCards := inSplit[1:]
	var cards []card

	for _, c := range rawCards {
		cards = append(cards, makeCard(c))
	}

drawloop:
	for _, draw := range draws {
		for _, bingoCard := range cards {
			bingoCard.check(draw)
			if bingoCard.winner() {
				winnerScore := bingoCard.score(draw)
				fmt.Printf("Winner found!\nScore: %d\n", winnerScore)
				break drawloop
			}
		}
	}
}
