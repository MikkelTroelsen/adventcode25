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
	currentNum := 50
	password := 0
	for _, rotation := range rotations {
		way := rotation[:1]
		clicks, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatal("Cant parse click to int: ", err)
		}
		fmt.Print("We move the dial: " + rotation)
		if way == "L" {
			currentNum += clicks
		}
		if way == "R" {
			currentNum -= clicks
		}
		//fmt.Print(" Current number: ")
		fmt.Println(currentNum)
		for currentNum < 0 {
			currentNum += 100
		}
		if currentNum > 100 {
			currentNum = 100 % currentNum
		}

		fmt.Print("Current number after mod: ")
		fmt.Println(currentNum)
		if currentNum == 0 {
			password++
		}
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
