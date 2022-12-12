package main

import (
	"fmt"
	"math"
	"sort"
	"util"
)

func main() {
	part12b()
}

// find the shortest path
// for each field, calculate the shortest valid possible path (from that field to the destination) and assign that value to that field
func part12a() {
	lines := util.ReadFile("12", "sample")
	//initialize distance-array, set all distances to max
	distances := make([][]int, len(lines))
	for i := range distances {
		distances[i] = make([]int, len(lines[0]))
		for distanceIndex := range distances[i] {
			distances[i][distanceIndex] = math.MaxInt
		}
	}

	//finding the position of the exit
	var finalX = 0
	var finalY = 0
	var startX = 0
	var startY = 0
	for y, line := range lines {
		for x, char := range line {
			if string(char) == "E" {
				distances[y][x] = 0
				//replacing the exit-symbol with '{' for easier ascii-comparison
				var helper = lines[y][:x]
				helper += "z"
				helper += lines[y][x+1:]
				lines[y] = helper
				finalX = x
				finalY = y
			}
			if string(char) == "S" {
				distances[y][x] = math.MaxInt
				//replacing the exit-symbol with '{' for easier ascii-comparison
				var helper = lines[y][:x]
				helper += "a"
				helper += lines[y][x+1:]
				lines[y] = helper
				startX = x
				startY = y
			}
		}
	}

	distances = findShortestWays(lines, distances, finalX, finalY, startX, startY)

	for _, row := range distances {
		fmt.Println(row)
	}
	fmt.Println(distances[startY][startX])
}

func part12b() {
	lines := util.ReadFile("12", "data")
	//initialize distance-array, set all distances to max
	distances := make([][]int, len(lines))
	for i := range distances {
		distances[i] = make([]int, len(lines[0]))
		for distanceIndex := range distances[i] {
			distances[i][distanceIndex] = math.MaxInt
		}
	}

	//finding the position of the exit
	var finalX = 0
	var finalY = 0
	var startX = 0
	var startY = 0
	for y, line := range lines {
		for x, char := range line {
			if string(char) == "E" {
				distances[y][x] = 0
				//replacing the exit-symbol with '{' for easier ascii-comparison
				var helper = lines[y][:x]
				helper += "z"
				helper += lines[y][x+1:]
				lines[y] = helper
				finalX = x
				finalY = y
			}
			if string(char) == "S" {
				distances[y][x] = math.MaxInt
				//replacing the exit-symbol with '{' for easier ascii-comparison
				var helper = lines[y][:x]
				helper += "a"
				helper += lines[y][x+1:]
				lines[y] = helper
				startX = x
				startY = y
			}
		}
	}

	distances = findShortestWays(lines, distances, finalX, finalY, startX, startY)
	startingPoints := make([]int, 0)
	for rI, row := range distances {
		for cI, _ := range row {
			if lines[rI][cI] == 'a' {
				startingPoints = append(startingPoints, distances[rI][cI])
			}
		}
	}
	sort.Ints(startingPoints)
	fmt.Println(startingPoints[0])
}

// for all valid steps calculate the distance
func findShortestWays(lines []string, distances [][]int, x int, y int, startX int, startY int) [][]int {
	//top
	if y > 0 {
		if lines[y-1][x] >= lines[y][x]-1 {
			if distances[y-1][x] > distances[y][x]+1 {
				distances[y-1][x] = distances[y][x] + 1
				distances = findShortestWays(lines, distances, x, y-1, startX, startY)
			}
		}
	}
	//bottom
	if y < len(lines)-1 {
		if lines[y+1][x] >= lines[y][x]-1 {
			if distances[y+1][x] > distances[y][x]+1 {
				distances[y+1][x] = distances[y][x] + 1
				distances = findShortestWays(lines, distances, x, y+1, startX, startY)
			}
		}
	}
	//left
	if x > 0 {
		if lines[y][x-1] >= lines[y][x]-1 {
			if distances[y][x-1] > distances[y][x]+1 {
				distances[y][x-1] = distances[y][x] + 1
				distances = findShortestWays(lines, distances, x-1, y, startX, startY)
			}
		}
	}
	//right
	if x < len(lines[0])-1 {
		if lines[y][x+1] >= lines[y][x]-1 {
			if distances[y][x+1] > distances[y][x]+1 {
				distances[y][x+1] = distances[y][x] + 1
				distances = findShortestWays(lines, distances, x+1, y, startX, startY)
			}
		}
	}
	//fmt.Println("after checking: ", x, ",", y, " [", string(lines[y][x]), "]")
	//for _, row := range distances {
	//	fmt.Println(row)
	//}
	return distances
}
