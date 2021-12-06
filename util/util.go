package util

import (
	"bufio"
	"io/ioutil"
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

func InputRaw() (string, error) {
	file, err := ioutil.ReadFile("./in.txt")
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func StrsToInts(strs []string) []int {
	var result []int
	for _, s := range strs {
		c, _ := strconv.Atoi(s)
		result = append(result, c)
	}
	return result
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
