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

func validGame(input string, constraints map[string]int) bool {
	a := strings.Split(input, ": ")[1]
	b := strings.Split(a, "; ")
	fmt.Println(b)
	for _, c := range b {
		d := strings.Split(c, " ")

		nr, err := strconv.Atoi(d[0])
		if err != nil {
			panic(err)
		}

		color := d[1][0 : len(d[1])-1]
		//fmt.Println(nr, color)
		if constraints[color] < nr {
			return false
		}
	}
	return true
}

func main() {

	a, err := readInput("inputt.txt")
	if err != nil {
		panic(err)
	}

	constraints := map[string]int{"red": 12, "green": 14, "blue": 13}

	inputs := strings.Split(a, "\n")[0 : len(strings.Split(a, "\n"))-1]
	for i, input := range inputs {
		fmt.Println(i+1, validGame(input, constraints), input)
	}
}
