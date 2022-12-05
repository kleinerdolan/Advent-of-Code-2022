package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

func main() {
	part05b()
}

// read stacks
// loop through instructions
// generate result-string
func part05a() {
	lines := util.ReadFileUntrimmed("05", "data")
	var startOfInstructions = 0
	var numberOfStacks = 0

	//read stacks
	for index, line := range lines {
		if !strings.Contains(line, "[") {
			startOfInstructions = index + 2
			//read the number of stacks
			numberOfStacks, _ = strconv.Atoi(string(lines[startOfInstructions-2][len(lines[startOfInstructions-2])-1]))
			break
		}
	}

	var stacks = make([][]string, numberOfStacks)
	for _, line := range lines {
		if !strings.Contains(line, "[") {
			break
		}
		for i := 0; i < len(line); i++ {
			if line[i] >= 65 && line[i] <= 90 {
				stacks[(i / 4)] = append(stacks[(i/4)], string(line[i]))
			}
		}
	}
	//reverse the stacks
	for i, stack := range stacks {
		stacks[i] = reverseStack(stack)
	}

	//parse instructions
	for i := startOfInstructions; i < len(lines); i++ {
		var numOfCrates, err1 = strconv.Atoi(strings.Split(strings.Split(lines[i], "move ")[1], " ")[0])
		if err1 != nil {
			log.Fatal("Atoi() failed")
		}
		var fromCrate, err2 = strconv.Atoi(strings.Split(strings.Split(lines[i], "from ")[1], " ")[0])
		if err2 != nil {
			log.Fatal("Atoi() failed")
		}
		var toCrate, err3 = strconv.Atoi(strings.Split(strings.Split(lines[i], "to ")[1], " ")[0])
		if err3 != nil {
			log.Fatal("Atoi() failed")
		}
		//zero-base
		fromCrate--
		toCrate--
		//execute instructions
		for j := 0; j < numOfCrates; j++ {
			//add to new stack
			stacks[toCrate] = append(stacks[toCrate], stacks[fromCrate][len(stacks[fromCrate])-1])
			//remove from old stack
			stacks[fromCrate] = stacks[fromCrate][:len(stacks[fromCrate])-1]
		}
	}
	fmt.Println(generateResultString(stacks))
}

func part05b() {
	lines := util.ReadFileUntrimmed("05", "data")
	var startOfInstructions = 0
	var numberOfStacks = 0

	//read stacks
	for index, line := range lines {
		if !strings.Contains(line, "[") {
			startOfInstructions = index + 2
			//read the number of stacks
			numberOfStacks, _ = strconv.Atoi(string(lines[startOfInstructions-2][len(lines[startOfInstructions-2])-1]))
			break
		}
	}

	var stacks = make([][]string, numberOfStacks)
	for _, line := range lines {
		if !strings.Contains(line, "[") {
			break
		}
		for i := 0; i < len(line); i++ {
			if line[i] >= 65 && line[i] <= 90 {
				stacks[(i / 4)] = append(stacks[(i/4)], string(line[i]))
			}
		}
	}
	//reverse the stacks
	for i, stack := range stacks {
		stacks[i] = reverseStack(stack)
	}

	//parse instructions
	for i := startOfInstructions; i < len(lines); i++ {
		var numOfCrates, err1 = strconv.Atoi(strings.Split(strings.Split(lines[i], "move ")[1], " ")[0])
		if err1 != nil {
			log.Fatal("Atoi() failed")
		}
		var fromCrate, err2 = strconv.Atoi(strings.Split(strings.Split(lines[i], "from ")[1], " ")[0])
		if err2 != nil {
			log.Fatal("Atoi() failed")
		}
		var toCrate, err3 = strconv.Atoi(strings.Split(strings.Split(lines[i], "to ")[1], " ")[0])
		if err3 != nil {
			log.Fatal("Atoi() failed")
		}
		//zero-base
		fromCrate--
		toCrate--
		//execute instructions
		//add to new stack
		stacks[toCrate] = append(stacks[toCrate], stacks[fromCrate][len(stacks[fromCrate])-numOfCrates:]...)
		//remove from old stack
		stacks[fromCrate] = stacks[fromCrate][:len(stacks[fromCrate])-numOfCrates]

	}
	fmt.Println(generateResultString(stacks))
}

func reverseStack(stack []string) []string {
	var reversedStack []string
	for i := 0; i < len(stack); i++ {
		reversedStack = append(reversedStack, stack[len(stack)-1-i])
	}
	return reversedStack
}

func generateResultString(stacks [][]string) string {
	var result = ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}
	return result
}
