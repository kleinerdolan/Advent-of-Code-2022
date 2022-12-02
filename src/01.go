package main

import (
	"fmt"
	"log"
	"strconv"
	"util"
)

func main() {
	part01b()
}

func part01a() {
	lines := util.ReadFile("01", "data")
	var elfs [][]string
	var curElf []string
	var max = 0
	var curSum = 0
	for _, element := range lines {
		// index is the index where we are
		// element is the element from someSlice for where we are
		if len(element) == 0 {
			curSum = sum(curElf)
			if curSum > max {
				max = curSum
			}
			elfs = append(elfs, curElf)
			curElf = nil
		} else {
			curElf = append(curElf, element)
		}
	}
	//appending the last elf
	elfs = append(elfs, curElf)
	fmt.Println(max)
}

func part01b() {
	lines := util.ReadFile("01", "data")
	var elfs [][]string
	var curElf []string
	var max1 = 0
	var max2 = 0
	var max3 = 0
	var curSum = 0
	for index, element := range lines {
		if len(element) == 0 {
			curSum = sum(curElf)
			updateTop3(curSum, &max1, &max2, &max3)
			elfs = append(elfs, curElf)
			curElf = nil
		} else {
			curElf = append(curElf, element)
			if index == len(lines)-1 {
				curSum = sum(curElf)
				updateTop3(curSum, &max1, &max2, &max3)
			}
		}
	}
	fmt.Println(max1 + max2 + max3)
}

// sums up numbers in a string-array (["1", "5", "3"] => 9)
func sum(slice []string) int {
	result := 0
	for _, value := range slice {
		curNum, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		result += curNum
	}
	return result
}

func updateTop3(curSum int, max1 *int, max2 *int, max3 *int) {
	if curSum > *max1 {
		*max3 = *max2
		*max2 = *max1
		*max1 = curSum
		return
	}
	if curSum > *max2 {
		*max3 = *max2
		*max2 = curSum
		return
	}
	if curSum > *max3 {
		*max3 = curSum
		return
	}
}
