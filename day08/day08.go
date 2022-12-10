package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix       = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
func viewScore(col []int, row []int, colPos int, rowPos int) (score int) {

	var (
		lView int = 0
		rView int = 0
		uView int = 0
		dView int = 0
	)
	if (colPos == 0) || (rowPos == 0) || (colPos == len(col) || (rowPos == len(row))) {
		score = 0
	} else {
		for x := rowPos - 1; x >= 0; x-- {
			lView++
			if row[rowPos] <= row[x] {
				break
			}
		}
		for x := rowPos + 1; x < len(row); x++ {
			rView++
			if row[rowPos] <= row[x] {
				break
			}
		}
		for y := colPos - 1; y >= 0; y-- {
			dView++
			if col[colPos] <= col[y] {
				break
			}
		}
		for y := colPos + 1; y < len(col); y++ {
			uView++
			if col[colPos] <= col[y] {
				break
			}
		}
	}
	score = uView * dView * lView * rView
	return score
}

func isVisible(col []int, row []int, colPos int, rowPos int) (visible bool) {
	visible = true
	colDown := true
	rowDown := true
	colUp := true
	rowUp := true
	//curTree := col[colPos]
	//fmt.Println(curTree)
	if (colPos == 0) || (rowPos == 0) || (colPos == len(col) || (rowPos == len(row))) {
		visible = true
	} else {
		for x := rowPos - 1; x >= 0; x-- {
			if row[rowPos] <= row[x] {
				rowDown = false
				break
			}
		}
		for x := rowPos + 1; x < len(row); x++ {
			if row[rowPos] <= row[x] {
				rowUp = false
				break
			}
		}
		for y := colPos - 1; y >= 0; y-- {
			if col[colPos] <= col[y] {
				colDown = false
				break
			}
		}
		for y := colPos + 1; y < len(col); y++ {
			if col[colPos] <= col[y] {
				colUp = false
				break
			}
		}
	}
	visible = colUp || colDown || rowUp || rowDown
	return visible
}

func main() {

	f, err := os.Open("/home/relliott/aoc_2022/day08/input_day08.txt")
	check(err)
	r := bufio.NewReader(f)
	var aocInput []string
	var s, e = Readln(r)
	aocInput = append(aocInput, s)
	for e == nil {
		s, e = Readln(r)
		aocInput = append(aocInput, s)
	}
	f.Close()

	var treeRows [99][]int
	var treeCols [99][]int
	var topViewScore int = 0
	for x := 0; x < len(aocInput); x++ {
		for y := 0; y < len(aocInput[x]); y++ {
			var tmp int
			tmp, e = strconv.Atoi(string(aocInput[x][y]))
			check(e)
			treeRows[x] = append(treeRows[x], tmp)
		}
	}
	var numVisible int = 0
	for rowPos := 0; rowPos < 99; rowPos++ {
		for colPos := 0; colPos < 99; colPos++ {
			treeCols[rowPos] = append(treeCols[rowPos], treeRows[colPos][rowPos])
		}
	}
	//fmt.Println(treeRows)
	//fmt.Println(treeCols)
	for row := 0; row < len(treeRows); row++ {
		for col := 0; col < len(treeCols); col++ {
			if isVisible(treeCols[row], treeRows[col], col, row) {
				score := viewScore(treeCols[row], treeRows[col], col, row)
				if topViewScore < score {
					topViewScore = score
				}
				numVisible++
			}
		}
	}

	fmt.Println("Number visible: ", numVisible)
	fmt.Println("Top view Score: ", topViewScore)
	//	yRow := trees[3]
	//	var xCol []int
	//	for y := 0; y < len(trees); y++ {
	//		xCol = append(xCol, trees[y][3])
	//	}
	//	fmt.Println(yRow)
	//	fmt.Println(xCol)

	//	visible := isVisible(xCol, yRow, 3, 3)

	//fmt.Println(visible)

}

// Guesses P1 4142 too high
