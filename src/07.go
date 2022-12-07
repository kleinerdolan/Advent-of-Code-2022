package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"util"
)

func main() {
	part07b()
}

// read each line, always keep track of the position
// save all found directories in a map = directory structure
// save content of a directory in the map. If it is another directory, link to that map-entry
func part07a() {
	var curPath = []string{"/"}
	var directoryMap = make(map[string]int)
	directoryMap["/"] = 0
	lines := util.ReadFile("07", "data")
	for _, line := range lines {

		//if it contains 'ls' all files are listed in the following lines - can be ignored
		if strings.Contains(line, "$ ls") {
			continue
		}
		//if it contains '$ cd', the directory is changed
		if strings.Contains(line, "$ cd") {
			if strings.Contains(line, "cd /") {
				curPath = []string{"/"}
			} else if strings.Contains(line, "..") {
				curPath = curPath[:len(curPath)-1]
			} else {
				curPath = append(curPath, strings.Split(line, "$ cd ")[1])
			}
			continue
		}
		//directory found , create new map-entry if it does not yet exist
		if strings.Contains(line, "dir ") {
			var tempDirectory = append(curPath, strings.Split(line, "dir ")[1])
			var directoryString = ""
			//workaround: create string representation of the current working directory
			for _, element := range tempDirectory {
				if directoryString != "" {
					directoryString += ","
				}
				directoryString += element
			}
			if _, exists := directoryMap[directoryString]; !exists {
				//fmt.Println("new directory: ", directoryString)
				directoryMap[directoryString] = 0
			}
			continue
		}
		//else: file found
		//workaround: create string representation of the current working directory
		var directoryString = ""
		for _, element := range curPath {
			if directoryString != "" {
				directoryString += ","
			}
			directoryString += element
		}
		for directoryString != "" {
			var currentSize, err = strconv.Atoi(strings.Split(line, " ")[0])
			if err != nil {
				log.Fatal("Parsing int failed")
			}
			directoryMap[directoryString] += currentSize
			if directoryString == "/" {
				directoryString = ""
				continue
			}
			directoryString = directoryString[:strings.LastIndex(directoryString, ",")]
		}
	}
	var result = 0
	//calculate result
	for _, element := range directoryMap {
		if element <= 100000 {
			result += element
		}
	}
	fmt.Println(result)
}

func part07b() {
	var curPath = []string{"/"}
	var directoryMap = make(map[string]int)
	directoryMap["/"] = 0
	lines := util.ReadFile("07", "data")
	for _, line := range lines {

		//if it contains 'ls' all files are listed in the following lines - can be ignored
		if strings.Contains(line, "$ ls") {
			continue
		}
		//if it contains '$ cd', the directory is changed
		if strings.Contains(line, "$ cd") {
			if strings.Contains(line, "cd /") {
				curPath = []string{"/"}
			} else if strings.Contains(line, "..") {
				curPath = curPath[:len(curPath)-1]
			} else {
				curPath = append(curPath, strings.Split(line, "$ cd ")[1])
			}
			continue
		}
		//directory found , create new map-entry if it does not yet exist
		if strings.Contains(line, "dir ") {
			var tempDirectory = append(curPath, strings.Split(line, "dir ")[1])
			var directoryString = ""
			//workaround: create string representation of the current working directory
			for _, element := range tempDirectory {
				if directoryString != "" {
					directoryString += ","
				}
				directoryString += element
			}
			if _, exists := directoryMap[directoryString]; !exists {
				//fmt.Println("new directory: ", directoryString)
				directoryMap[directoryString] = 0
			}
			continue
		}
		//else: file found
		//workaround: create string representation of the current working directory
		var directoryString = ""
		for _, element := range curPath {
			if directoryString != "" {
				directoryString += ","
			}
			directoryString += element
		}
		for directoryString != "" {
			var currentSize, err = strconv.Atoi(strings.Split(line, " ")[0])
			if err != nil {
				log.Fatal("Parsing int failed")
			}
			directoryMap[directoryString] += currentSize
			if directoryString == "/" {
				directoryString = ""
				continue
			}
			directoryString = directoryString[:strings.LastIndex(directoryString, ",")]
		}
	}
	//calculate needed space
	var unusedSpace = 70000000 - directoryMap["/"]
	fmt.Println(unusedSpace)
	var neededSpace = 30000000 - unusedSpace
	fmt.Println(neededSpace)
	var currentCandidateSpace = math.MaxInt
	var currentCandidateName = ""
	for key, element := range directoryMap {
		if element >= neededSpace {
			//check if is smaller than the previous candidate
			if element < currentCandidateSpace {
				currentCandidateSpace = element
				currentCandidateName = key
			}
		}
	}
	fmt.Println(currentCandidateName, ":", currentCandidateSpace)
}
