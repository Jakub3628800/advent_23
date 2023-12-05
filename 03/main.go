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

func isSymbol(c byte) bool {
	return c == ' ' || c == ',' || c == ':' || c == ';'
}

func isSymbolNear(i int, j int, inputarr *[]string) bool {
	return false
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}
	inputarr := strings.Split(input, "\n")

	processingNumber := 0
	processingNumberSymbolAround := false

	for i, line := range inputarr {
		for j, c := range line {
			if unicode.IsDigit(c) {
				processingNumber = processingNumber*10 + int(c) - 48
				if isSymbolNear(i, j, &inputarr) {
					processingNumberSymbolAround = true
				}

			} else {
				fmt.Println(processingNumber, processingNumberSymbolAround)
				processingNumber = 0
			}

		}

	}
}
