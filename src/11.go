package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

func main() {
	part11a()
}

// read monkeys:
// -array of arras containing items
// -array of operation: + old , * 3
// -array of test: 5,5,7 (divisible by 5 to 5 else 7)
// play 20 rounds
func part11a() {
	var monkeyIndex = 0
	monkeyItems := make([][]int, 8)
	monkeyOperation := make([]string, 8)
	monkeyTest := make([][]int, 8)
	lines := util.ReadFile("11", "data")
	//parsing the input
	for _, line := range lines {
		//parsing current monkey is finished
		if line == "" {
			monkeyIndex++
		}
		//parsing starting items
		if strings.Contains(line, "Starting items") {
			var itemsAsStrings = strings.Split(strings.Split(line, ": ")[1], " ")
			for _, itemString := range itemsAsStrings {
				var convertedItem, err = strconv.Atoi(strings.Split(itemString, ",")[0])
				if err != nil {
					log.Fatal("Converting items failed")
				}
				monkeyItems[monkeyIndex] = append(monkeyItems[monkeyIndex], convertedItem)
			}
		}
		//parsing operation
		if strings.Contains(line, "Operation") {
			monkeyOperation[monkeyIndex] = strings.Split(line, "new = old ")[1]
		}
		//parsing test
		if strings.Contains(line, "Test:") {
			var divider, err = strconv.Atoi(strings.Split(line, "divisible by ")[1])
			if err != nil {
				log.Fatal("Converting divider failed")
			}
			monkeyTest[monkeyIndex] = []int{divider}
		}
		if strings.Contains(line, "true:") {
			var destination, err = strconv.Atoi(strings.Split(line, "throw to monkey ")[1])
			if err != nil {
				log.Fatal("Converting destination failed")
			}
			monkeyTest[monkeyIndex] = append(monkeyTest[monkeyIndex], destination)
		}
		if strings.Contains(line, "false:") {
			var destination, err = strconv.Atoi(strings.Split(line, "throw to monkey ")[1])
			if err != nil {
				log.Fatal("Converting destination failed")
			}
			monkeyTest[monkeyIndex] = append(monkeyTest[monkeyIndex], destination)
		}
	}

	//for 20 times, iterate through each monkeys items and process/throw them
	//count the actions of each monkey
	var monkeyActions = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for round := 0; round < 20; round++ {
		for monkey := 0; monkey < 8; monkey++ {
			for item := 0; item < len(monkeyItems[monkey]); item++ {
				//apply operation for item
				monkeyActions[monkey]++
				if strings.Contains(monkeyOperation[monkey], "*") {
					if strings.Contains(monkeyOperation[monkey], "old") {
						monkeyItems[monkey][item] *= monkeyItems[monkey][item]
					} else {
						var operator, err = strconv.Atoi(strings.Split(monkeyOperation[monkey], " ")[1])
						if err != nil {
							log.Fatal("Converting operator failed")
						}
						monkeyItems[monkey][item] *= operator
					}
				} else {
					if strings.Contains(monkeyOperation[monkey], "old") {
						monkeyItems[monkey][item] += monkeyItems[monkey][item]
					} else {
						var operator, err = strconv.Atoi(strings.Split(monkeyOperation[monkey], " ")[1])
						if err != nil {
							log.Fatal("Converting operator failed")
						}
						monkeyItems[monkey][item] += operator
					}
				}
				//monkey is bored
				monkeyItems[monkey][item] /= 3

				//check divisible
				if monkeyItems[monkey][item]%monkeyTest[monkey][0] == 0 {
					monkeyItems[monkeyTest[monkey][1]] = append(monkeyItems[monkeyTest[monkey][1]], monkeyItems[monkey][item])
				} else {
					monkeyItems[monkeyTest[monkey][2]] = append(monkeyItems[monkeyTest[monkey][2]], monkeyItems[monkey][item])
				}
			}
			monkeyItems[monkey] = make([]int, 0)
		}
		//print monkeys
		//fmt.Println("After round ", round+1)
		//for i, monkey := range monkeyItems {
		//	fmt.Println("Monkey ", i, ": ", monkey)
		//}
	}
	fmt.Println(monkeyActions)
}
