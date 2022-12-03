package main

import (
	"fmt"
	"strings"
	"util"
)

func main() {
	part03b()
}

// split string in half
// search for same letter
// calculate value of letter
func part03a() {
	var totalsum = 0
	var currentA = ""
	var currentB = ""
	lines := util.ReadFile("03", "data")
	for _, line := range lines {
		currentA = line[:len(line)/2]
		currentB = line[len(line)/2:]
		for _, char := range currentA {
			if strings.Contains(currentB, string(char)) {
				totalsum += evaluateLetter(int(char))
				break
			}
		}
	}
	fmt.Println(totalsum)
}

func evaluateLetter(letter int) int {
	var value = 0
	if letter > 96 {
		value = letter - 96
	}
	if letter < 96 {
		value = letter - 38
	}
	return value
}

// group 3 lines together
// search for the single letter appearing in all three lines
// evaluate letter
func part03b() {
	var totalsum = 0
	var currentA = ""
	var currentB = ""
	var currentC = ""
	lines := util.ReadFile("03", "data")
	for i := 0; i <= len(lines)-1; i += 3 {
		currentA = lines[i]
		currentB = lines[i+1]
		currentC = lines[i+2]
		for _, char := range currentA {
			if strings.Contains(currentB, string(char)) {
				if strings.Contains(currentC, string(char)) {
					totalsum += evaluateLetter(int(char))
					break
				}
			}
		}
	}
	fmt.Println(totalsum)
}
