package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(filepath string) (string, error) {

	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(content), nil

}

func intersection(a, b []int) []int {
	var result []int

	for _, aNr := range a {
		for _, bNr := range b {
			if aNr == bNr {
				result = append(result, aNr)
			}
		}
	}

	return result
}

func lineToSlice(line string) []int {

	var result []int

	for _, number := range strings.Split(line, " ") {
		nr, _ := strconv.Atoi(number)
		if nr != 0 {
			result = append(result, nr)
		}
	}

	return result
}

func cardScore(card string) float64 {
	removedPrefix := strings.Split(strings.Split(card, ":")[1], "|")

	winningNumbers := lineToSlice(removedPrefix[0])
	ourNumbers := lineToSlice(removedPrefix[1])
	matchNr := len(intersection(winningNumbers, ourNumbers))
	if matchNr == 0 {
		return 0
	}
	return math.Pow(2, float64(matchNr-1))
}

func main() {

	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	result := float64(0)
	for _, line := range strings.Split(input, "\n")[0 : len(strings.Split(input, "\n"))-1] {
		result += cardScore(line)
	}
	fmt.Println(result)
}
