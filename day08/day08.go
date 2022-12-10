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

func isVisible(col []int, row []int, colPos int, rowPos int) (visible bool) {
	visible = true
	colDown := true
	rowDown := true
	colUp := true
	rowUp := true

	if col[colPos] != row[rowPos] {
		fmt.Println("Not looking at same value")
	}

	if (colPos == 0) || (rowPos == 0) || (colPos == len(col) || (rowPos == len(row))) {
		visible = true
	} else {
		for x := rowPos - 1; x >= 0; x-- {
			if row[rowPos] < row[x] {
				rowDown = false
				break
			}
		}
		for x := rowPos + 1; x < len(row); x++ {
			if row[rowPos] < row[x] {
				rowUp = false
				break
			}
		}
		for y := colPos - 1; y >= 0; y-- {
			if col[colPos] < col[y] {
				colDown = false
				break
			}
		}
		for y := colPos + 1; y < len(col); y++ {
			if col[colPos] < col[y] {
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

	var trees [5][]int

	for x := 0; x < len(aocInput); x++ {
		for y := 0; y < len(aocInput[x]); y++ {
			var tmp int
			tmp, e = strconv.Atoi(string(aocInput[x][y]))
			check(e)
			trees[x] = append(trees[x], tmp)
		}
	}
	//fmt.Println(trees)
	var row []int
	var col []int
	var numVisible int = 0
	for rowPos := 0; rowPos < 5; rowPos++ {
		col = nil
		row = trees[rowPos]
		for colPos := 0; colPos < 5; colPos++ {
			col = append(col, trees[colPos][rowPos])
		}
		fmt.Println("col:", col)
		fmt.Println("Row: ", row)

	}
	numVisible = numVisible
	fmt.Println("Number visible: ", numVisible)
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
