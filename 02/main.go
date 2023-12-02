package main

import (
	"fmt"
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

func validGame(input string, constraints map[string]int) int {

	gameSplit := strings.Split(input, ": ")
	gameTitle := gameSplit[0]
	gameTurns := gameSplit[1]

	for _, gameTurn := range strings.Split(gameTurns, "; ") {

		for _, cubeColorNumber := range strings.Split(gameTurn, ", ") {

			cubeColorNr := strings.Split(cubeColorNumber, " ")

			nr, err := strconv.Atoi(cubeColorNr[0])
			if err != nil {
				panic(err)
			}
			if constraints[cubeColorNr[1]] < nr {
				return 0
			}
		}

	}
	aa, err := strconv.Atoi(strings.Split(gameTitle, " ")[1])
	if err != nil {
		panic(err)
	}
	return aa
}

func main() {

	a, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	constraints := map[string]int{"red": 12, "green": 13, "blue": 14}

	inputs := strings.Split(a, "\n")[0 : len(strings.Split(a, "\n"))-1]
	result := 0
	for _, input := range inputs {
		result += validGame(input, constraints)
		//fmt.Println(validGame(input, constraints), input)
	}

	fmt.Println(result)
}
