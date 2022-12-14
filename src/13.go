package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

func main() {
	part13a()
}

func part13a() {
	var pairIndex = 0
	lines := util.ReadFile("13", "data")
	pairs := make([][]string, len(lines)/3+1)
	//saving input in an array
	for _, line := range lines {
		if len(line) == 0 {
			pairIndex++
			continue
		}
		if len(pairs[pairIndex]) == 0 {
			pairs[pairIndex] = append(pairs[pairIndex], line)
		} else {
			pairs[pairIndex] = append(pairs[pairIndex], line)
		}
	}
	var correctIndexesSum = 0
	var correctPairs = 0
	for i, pair := range pairs {
		if checkPairIsValid(pair) {
			correctPairs++
			correctIndexesSum += i + 1
		}
	}
	fmt.Println(correctPairs)
	fmt.Println(correctIndexesSum)
}

// iterate through left numbers
// for each left number, compare to the next right number (keep track of already checked numbers)
// if right has no more numbers -> return false
// if right is smaller than left -> return false
func checkPairIsValid(pair []string) bool {

	//comparisons involving '[]' are missing!

	f := func(c rune) bool {
		if c == '[' {
			return true
		}
		if c == ']' {
			return true
		}
		if c == ',' {
			return true
		}
		return false
	}
	//find next number
	var splitPair0String = strings.FieldsFunc(pair[0], f)
	var splitPair1String = strings.FieldsFunc(pair[1], f)

	splitPair0 := make([]int, len(splitPair0String))
	splitPair1 := make([]int, len(splitPair1String))

	for i, s := range splitPair0String {
		var convertedNumber, err = strconv.Atoi(s)
		if err != nil {
			log.Fatal("number conversion failed")
		}
		splitPair0[i] = convertedNumber
	}

	for i, s := range splitPair1String {
		var convertedNumber, err = strconv.Atoi(s)
		if err != nil {
			log.Fatal("number conversion failed")
		}
		splitPair1[i] = convertedNumber
	}

	fmt.Println(pair)
	fmt.Println(splitPair0, " ", splitPair1, " ")

	if len(splitPair0) == 0 && len(splitPair1) == 0 {
		fmt.Println("false both empty")
		return false
	}

	var biggerListLength = 0
	if len(splitPair0) >= len(splitPair1) {
		biggerListLength = len(splitPair0)
	} else {
		biggerListLength = len(splitPair1)
	}
	for i := 0; i < biggerListLength; i++ {
		//check size
		if len(splitPair0) == i {
			fmt.Println("true Left side ran out of items")
			return true
		}
		if len(splitPair1) == i {
			fmt.Println("false Right side ran out of items")
			return false
		}
		fmt.Println("comparing ", splitPair0[i], " and ", splitPair1[i])
		if splitPair0[i] < splitPair1[i] {
			fmt.Print(splitPair0[i], " < ", splitPair1[i])
			fmt.Println(" true left side is smaller")
			return true
		}
		if splitPair0[i] > splitPair1[i] {
			fmt.Print(splitPair0[i], " > ", splitPair1[i])
			fmt.Println(" false right side is smaller")
			return false
		}
	}
	fmt.Println("true")
	return true
}
