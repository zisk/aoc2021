package util

import (
	"bufio"
	"os"
	"strconv"
)

func InputToInts() ([]int, error) {
	var result []int

	file, err := os.Open("./in.txt")
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		linen, _ := strconv.Atoi(line)
		result = append(result, linen)
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil

}

func InputToTxt() ([]string, error) {
	var result []string

	file, err := os.Open("./in.txt")
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}
