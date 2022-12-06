package main

import (
	"fmt"
	"util"
)

func main() {
	part06b()
}

// read chars one by one
// check if there is a duplicate in the last 4 chars
func part06a() {
	lines := util.ReadFile("06", "data")
	for i := 3; i < len(lines[0]); {
		if lines[0][i] == lines[0][i-1] {
			i += 3
			continue
		} else if lines[0][i] == lines[0][i-2] || lines[0][i-1] == lines[0][i-2] {
			i += 2
			continue
		} else if lines[0][i] == lines[0][i-3] || lines[0][i-1] == lines[0][i-3] || lines[0][i-2] == lines[0][i-3] {
			i += 1
			continue
		}
		//zero-based!
		fmt.Println(i + 1)
		break
	}
}

// create a set
func part06b() {
	lines := util.ReadFile("06", "data")
	for i := 0; i < len(lines[0]); {
		var set = make(map[string]int)
		for j := i; j < i+15; j++ {
			if len(set) == 14 {
				//zero-based!
				fmt.Println(i + 14)
				return
			}
			if _, ok := set[string(lines[0][j])]; ok {
				i = set[string(lines[0][j])]
				break
			} else {
				set[string(lines[0][j])] = j
			}
		}
		i++
	}
}
