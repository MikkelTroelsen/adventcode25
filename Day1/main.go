package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	rotations := parseInput()
	// rotations := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	part1(&rotations)
	part2(&rotations)
}

func part1(rotations *[]string) {
	currentNum := 50
	password := 0
	for _, rotation := range *rotations {
		way := rotation[:1]
		clicks, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatal("Cant parse click to int: ", err)
		}
		if way == "L" {
			currentNum -= clicks
		} else {
			currentNum += clicks
		}
		if currentNum%100 == 0 {
			password++
		}
	}
	fmt.Println(password)
}

func part2(rotations *[]string) {
	currentNum := 50
	password := 0
	fmt.Println(currentNum, "Password: ", password)

	prev := 50
	for _, rotation := range *rotations {

		way := rotation[:1]
		clicks, err := strconv.Atoi(rotation[1:])

		if err != nil {
			log.Fatal("Cant parse click to int: ", err)
		}
		if way == "L" {
			currentNum -= clicks
		} else {
			currentNum += clicks
		}
		if way == "L" && prev == 0 {
			fmt.Println("prev", prev)
			password--
		}

		for currentNum > 100 {
			currentNum -= 100
			password++
		}
		for currentNum < 0 {
			currentNum += 100
			password++
		}
		if currentNum == 0 {
			password++
		}
		if currentNum == 100 {
			currentNum -= 100
			password++
		}

		prev = currentNum

		fmt.Println(rotation, " -> ", currentNum, "Password: ", password)
	}
	fmt.Println(password)
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
