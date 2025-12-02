package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput()
	//input := "8334-13124"
	//input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	part1(input)
}

func part1(input string) {
	sequances := strings.Split(input, ",")
	combined := 0
	for _, sequance := range sequances {
		splitSequance := strings.Split(sequance, "-")
		leftRange := splitSequance[0]
		rightRange := splitSequance[1]
		leftLen := len(leftRange)
		rightLen := len(rightRange)
		if leftLen%2 != 0 {
			leftRange = strconv.FormatFloat(
				math.Pow(10, float64(leftLen)),
				'f', // formant
				-1,  // precision -1 for automatic
				64,  // float64
			)
			leftLen++
		}
		if rightLen%2 != 0 {
			rightRange = strconv.FormatFloat(
				math.Pow(10, float64(rightLen-1))-1,
				'f', // format
				-1,  // precision
				64,  // float64
			)
			rightLen--
		}
		leftMid := leftLen / 2
		ll, _ := strconv.Atoi(leftRange[:leftMid])
		lr, _ := strconv.Atoi(leftRange[leftMid:])

		rightMid := rightLen / 2
		rl, _ := strconv.Atoi(rightRange[:rightMid])
		rr, _ := strconv.Atoi(rightRange[rightMid:])

		if ll < lr {
			ll++
		}
		if rl > rr {
			rl--
		}
		for i := ll; i <= rl; i++ {
			dup := DuplicateDigits(i)
			combined += dup
		}

	}
	println(combined)
}

func DuplicateDigits(digit int) int {
	strDigit := strconv.Itoa(digit)
	doubleDigit, _ := strconv.Atoi(strDigit + strDigit)
	return doubleDigit
}

func parseInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	result := strings.TrimSpace(string(data))
	return result
}
