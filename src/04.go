package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

func main() {
	part04b()
}

// split string at comma
// check a starts and ends inside b or the opposite
func part04a() {
	lines := util.ReadFile("04", "data")
	var totalsum = 0
	var curLineSplit []string
	for _, line := range lines {
		curLineSplit = strings.Split(line, ",")
		if checkAContainsB(curLineSplit[0], curLineSplit[1]) || checkAContainsB(curLineSplit[1], curLineSplit[0]) {
			totalsum += 1
		}
	}
	fmt.Println(totalsum)
}

func checkAContainsB(a string, b string) bool {
	var startA, err1 = strconv.Atoi(strings.Split(a, "-")[0])
	if err1 != nil {
		log.Fatal(err1)
	}
	var EndA, err2 = strconv.Atoi(strings.Split(a, "-")[1])
	if err2 != nil {
		log.Fatal(err2)
	}
	var startB, err3 = strconv.Atoi(strings.Split(b, "-")[0])
	if err3 != nil {
		log.Fatal(err3)
	}
	var EndB, err4 = strconv.Atoi(strings.Split(b, "-")[1])
	if err4 != nil {
		log.Fatal(err4)
	}

	//check b starts before a
	if startB < startA {
		return false
	}
	//check b ends after a
	if EndB > EndA {
		return false
	}
	//else: b is inside/same as a
	return true
}

func part04b() {
	lines := util.ReadFile("04", "data")
	var totalsum = 0
	var curLineSplit []string
	for _, line := range lines {
		curLineSplit = strings.Split(line, ",")
		if checkOverlap(curLineSplit[0], curLineSplit[1]) {
			totalsum += 1
		}
	}
	fmt.Println(totalsum)
}

func checkOverlap(a string, b string) bool {
	var startA, err1 = strconv.Atoi(strings.Split(a, "-")[0])
	if err1 != nil {
		log.Fatal(err1)
	}
	var EndA, err2 = strconv.Atoi(strings.Split(a, "-")[1])
	if err2 != nil {
		log.Fatal(err2)
	}
	var startB, err3 = strconv.Atoi(strings.Split(b, "-")[0])
	if err3 != nil {
		log.Fatal(err3)
	}
	var EndB, err4 = strconv.Atoi(strings.Split(b, "-")[1])
	if err4 != nil {
		log.Fatal(err4)
	}

	//check b contains a part of a
	if startB <= startA && EndB >= startA {
		return true
	}
	//check a contains a part of b
	if startA <= startB && EndA >= startB {
		return true
	}
	//else: no overlap
	return false
}
