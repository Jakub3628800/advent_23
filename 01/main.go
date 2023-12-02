package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func readInput(filepath string) (string, error) {

	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(content), nil

}

func calibrationValue(input string) int {
	first := -1
	last := 10
	for _, char := range input {
		if unicode.IsDigit(char) {
			if first == -1 {
				first = int(char) - 48
				last = int(char) - 48
			} else {
				last = int(char) - 48
			}
		}
	}

	return 10*first + last
}

func calibrationValuePartTwo(input string) int {
	first := -1
	last := 10
	for _, char := range input {
		if unicode.IsDigit(char) {
			if first == -1 {
				first = int(char) - 48
				last = int(char) - 48
			} else {
				last = int(char) - 48
			}
		}
	}
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	for i, numb := range numbers {
		if strings.HasPrefix(input, numb) {
			first = i + 1
		}
		if strings.HasSuffix(input, numb) {
			last = i + 1
		}
	}

	return 10*first + last
}

func main() {
	a, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputs := strings.Split(a, "\n")
	resultSum := 0
	resultSumPartTwo := 0
	for _, input := range inputs {
		resultSum += calibrationValue(input)
		resultSumPartTwo += calibrationValuePartTwo(input)
		fmt.Println(input, calibrationValuePartTwo(input))
	}
	fmt.Println(resultSum)
	fmt.Println(resultSumPartTwo)
}
