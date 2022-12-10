package main

import (
	"fmt"
	"util"
)

func main() {
	part08b()
}

// create empty array of the same size as input
// add borders
// iterate through columns and determine visible trees for both directions - mark them in the new array
// iterate through lines and determine visible trees for both directions - mark them in the new array
func part08a() {
	lines := util.ReadFile("08", "data")
	//create an array of the same sie as the input for marking the visible trees
	visibleTrees := make([][]bool, len(lines))
	for i := range visibleTrees {
		visibleTrees[i] = make([]bool, len(lines))
	}
	for rowIndex, row := range lines {
		for columnIndex, _ := range lines[rowIndex] {
			//fmt.Println("cur col: ", string(lines[columnIndex][rowIndex]))
			//mark borders
			if rowIndex == 0 || rowIndex == len(lines)-1 || columnIndex == 0 || columnIndex == len(row)-1 {
				visibleTrees[columnIndex][rowIndex] = true
			}
		}
		//checking rows left to right
		var curMaxRow1 = row[0]
		for i := 1; i < len(row)-1; i++ {
			if row[i] > curMaxRow1 {
				curMaxRow1 = row[i]
				visibleTrees[rowIndex][i] = true
			}
		}

		//checking rows right to left
		var curMaxRow2 = row[len(row)-1]
		for i := len(row) - 1; i > 0; i-- {
			if row[i] > curMaxRow2 {
				curMaxRow2 = row[i]
				visibleTrees[rowIndex][i] = true
			}
		}

		//checking columns top to bottom
		var curMaxCol1 = lines[0][rowIndex]
		for i := 1; i < len(lines[rowIndex])-1; i++ {
			if lines[i][rowIndex] > curMaxCol1 {
				curMaxCol1 = lines[i][rowIndex]
				visibleTrees[i][rowIndex] = true
			}
		}

		//checking columns bottom to top
		var curMaxCol2 = lines[len(row)-1][rowIndex]
		for i := len(row) - 1; i > 0; i-- {
			if lines[i][rowIndex] > curMaxCol2 {
				curMaxCol2 = lines[i][rowIndex]
				visibleTrees[i][rowIndex] = true
			}
		}
	}

	//calculate result
	var visibleCount = 0
	for _, row := range visibleTrees {
		for _, tree := range row {
			if tree {
				visibleCount++
			}
		}
	}
	fmt.Println(visibleCount)
}

func part08b() {
	lines := util.ReadFile("08", "data")
	//create an array of the same sie as the input for marking the visible trees
	treeScore := make([][]int, len(lines))
	for i := range treeScore {
		treeScore[i] = make([]int, len(lines))
	}
	for rowIndex, _ := range lines {
		for columnIndex, _ := range lines[rowIndex] {
			//mark borders
			treeScore[rowIndex][columnIndex] = calculateScore(lines, rowIndex, columnIndex)
		}
	}
	var maxScore = 0
	for _, row := range treeScore {
		for _, tree := range row {
			if tree > maxScore {
				maxScore = tree
			}
		}
	}
	fmt.Println(maxScore)

}

func calculateScore(lines []string, row int, column int) int {
	//checking rows left to right
	var rightScore = 0
	for i := column + 1; i < len(lines); i++ {
		rightScore++
		if lines[row][i] >= lines[row][column] {
			break
		}
	}

	//checking rows right to left
	var leftScore = 0
	for i := column - 1; i >= 0; i-- {
		leftScore++
		if lines[row][i] >= lines[row][column] {
			break
		}
	}

	//checking columns top to bottom
	var bottomScore = 0
	for i := row + 1; i < len(lines[row]); i++ {
		bottomScore++
		if lines[i][column] >= lines[row][column] {
			break
		}
	}

	//checking columns bottom to top
	var topScore = 0
	for i := row - 1; i >= 0; i-- {
		topScore++
		if lines[i][column] >= lines[row][column] {
			break
		}
	}

	return rightScore * leftScore * topScore * bottomScore
}
