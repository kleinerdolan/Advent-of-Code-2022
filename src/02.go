package main

import (
	"fmt"
	"strings"
	"util"
)

func main() {
	part02b()
}

func part02a() {
	var totalsum = 0
	lines := util.ReadFile("02", "data")
	for _, line := range lines {
		//bonus for shape selection
		if strings.Contains(line, "X") {
			totalsum += 1
		} else if strings.Contains(line, "Y") {
			totalsum += 2
		} else if strings.Contains(line, "Z") {
			totalsum += 3
		}

		//check win or draw
		if line == "A Y" || line == "B Z" || line == "C X" {
			totalsum += 6
		} else if line == "A X" || line == "B Y" || line == "C Z" {
			totalsum += 3
		}
	}
	fmt.Println(totalsum)
}

func part02b() {
	var totalsum = 0
	lines := util.ReadFile("02", "data")
	for _, line := range lines {
		//check for win or draw
		if strings.Contains(line, "Y") {
			totalsum += 3
		} else if strings.Contains(line, "Z") {
			totalsum += 6
		}

		//check win or draw
		if line == "A Y" || line == "B X" || line == "C Z" {
			totalsum += 1
		} else if line == "A Z" || line == "B Y" || line == "C X" {
			totalsum += 2
		} else if line == "A X" || line == "B Z" || line == "C Y" {
			totalsum += 3
		}
	}
	fmt.Println(totalsum)
}

//part1:
//A, X(1) rock
//B, Y(2) Paper
//C, Z(3) Scissors
//win 6	AY BZ CX
//draw 3	AX BY CZ
//lose 0	AZ BX CY

//part2:
//A X -> lose with scissor = 0(lose) + 3(scissor)
//A Y -> draw with rock = 3(draw) + 1(rock)
//A Z -> win with paper = 6(win) + 2(paper)

//B X -> lose with rock = 0(lose) + 1(rock)
//B Y -> draw with paper = 3(draw) + 2(paper)
//B Z -> win with scissor = 6(win) + 3(scissor)

//C X -> lose with paper = 0(lose) + 2(paper)
//C Y -> draw with scissor = 3(draw) + 3(scissor)
//C Z -> win with rock = 6(win) + 1(rock)
