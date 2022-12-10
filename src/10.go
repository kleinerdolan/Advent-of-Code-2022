package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

func main() {
	part10b()
}

// go through all lines
// update the register-value
func part10a() {
	var register = 1
	var signalStrength = 0
	lines := util.ReadFile("10", "data")
	var stack []string
	for i := 0; i < len(lines)+1000; i++ {
		//var curLine = "/"
		//if i < len(lines)-1 {
		//	curLine = lines[i]
		//}
		fmt.Println(stack)
		//fmt.Println("cycle ", i+1, "current line: ", curLine, " (register=", register, ")")

		if (i+1-20)%40 == 0 {
			signalStrength += (i + 1) * register
		}

		//add new instruction to stack
		if i < len(lines)-1 {
			if strings.Contains(lines[i], "add") {
				stack = append(stack, "buffer")
			}
			stack = append(stack, lines[i])
		}
		if len(stack) == 0 && i > len(lines) {
			fmt.Println("signalstrength", signalStrength)
			return
		}

		//execute stack if there is an instruction
		if strings.Contains(stack[0], "addx") {
			var amountString = strings.Split(stack[0], " ")[1]
			var amountInt, err = strconv.Atoi(amountString)
			if err != nil {
				log.Fatal("Parsing string failed")
			}
			register += amountInt
		}
		stack = stack[1:]
	}
}

// iterate through all cycles
// build line
// after every 40 cycles, draw line
func part10b() {
	var register = 1
	curLine := make([]string, 0)
	lines := util.ReadFile("10", "data")
	var stack []string
	for i := 0; i < len(lines)+1000; i++ {

		if i%40 == 0 {
			fmt.Println(curLine)
			curLine = nil
		}

		//add new instruction to stack
		if i < len(lines)-1 {
			if strings.Contains(lines[i], "add") {
				stack = append(stack, "buffer")
			}
			stack = append(stack, lines[i])
		}
		if len(stack) == 0 && i > len(lines) {
			fmt.Println(curLine)
			return
		}

		curLine = drawNextPixel(curLine, register, i%40)

		//execute stack if there is an instruction
		if strings.Contains(stack[0], "add") {
			var amountString = strings.Split(stack[0], " ")[1]
			var amountInt, err = strconv.Atoi(amountString)
			if err != nil {
				log.Fatal("Parsing string failed")
			}
			register += amountInt

		}
		stack = stack[1:]
	}
}

func drawNextPixel(curLine []string, register int, cycle int) []string {
	if cycle >= register-1 && cycle <= register+1 {
		curLine = append(curLine, "#")
	} else {
		curLine = append(curLine, ".")
	}
	//fmt.Println("cycle: ", cycle+1, " register: ", register, " curLine: ", curLine)
	return curLine
}
