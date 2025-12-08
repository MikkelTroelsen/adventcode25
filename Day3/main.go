package main

import (
	"bufio"
	"os"
)

type partx func(string, *int)

func main() {
	// input := parseInput()
	input := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	loop(input, part1)
}

func loop(lines []string, logic partx) {
	result := 0

	for _, line := range lines {
		logic(line, &result)
	}

	println("---------------------------------")
	println(result)
}

func part1(str string, combined *int) {
	tens := 0
	ones := 0
	for _, r := range str {
		// TODO handle edge case last number is a bigger tens (last number should me able to change tens)
		number := int(r - '0')
		if number > tens {
			tens = number
			ones = 0
		} else if number > ones {
			ones = number
		}
	}
}

func parseInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}
