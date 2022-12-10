package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func moveBoxes(numBoxes int, startCol string, endCol string) (newStartCol string, newEndCol string) {
	newStartCol = startCol[:len(startCol)-numBoxes]
	tempStr := reverse(startCol[(len(startCol) - numBoxes):])
	newEndCol = endCol + tempStr

	return newStartCol, newEndCol
}

func moveBoxes2(numBoxes int, startCol string, endCol string) (newStartCol string, newEndCol string) {
	newStartCol = startCol[:len(startCol)-numBoxes]
	tempStr := startCol[(len(startCol) - numBoxes):]

	var movedStack string = ""
	var i int = 1
	numMoves := numBoxes / 3

	for numMoves != 0 {
		movedStack += tempStr[(len(tempStr) - 3*i):]
		tempStr = tempStr[:len(tempStr)-3]
		numMoves--
	}
	if numBoxes%3 != 0 {
		movedStack += tempStr
	}

	newEndCol += endCol + movedStack

	return newStartCol, newEndCol
}

func moveBoxes3(numBoxes int, startCol string, endCol string) (newStartCol string, newEndCol string) {
	newStartCol = startCol[:len(startCol)-numBoxes]
	tempStr := startCol[(len(startCol) - numBoxes):]
	newEndCol = endCol + tempStr

	return newStartCol, newEndCol
}

func getMoveOrder(s string) (numBoxes int, startCol int, endCol int, err error) {

	str1 := strings.SplitAfter(s, " ")

	numBoxes, err = strconv.Atoi(strings.TrimSpace(str1[1]))
	startCol, err = strconv.Atoi(strings.TrimSpace(str1[3]))
	endCol, err = strconv.Atoi(strings.TrimSpace(str1[5]))

	return numBoxes, startCol, endCol, err

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

func main() {

	f, err := os.Open("/home/relliott/aoc_2022/day05/input_day05.txt")
	check(err)
	r := bufio.NewReader(f)

	var (
		col1 string = "RGJBTVZ"
		col2 string = "JRVL"
		col3 string = "SQF"
		col4 string = "ZHNLFVQG"
		col5 string = "RQTJCSMW"
		col6 string = "SWTCHF"
		col7 string = "DZCVFNJ"
		col8 string = "LGZDWRFQ"
		col9 string = "JBWVP"
	)
	var cols = [9]string{col1, col2, col3, col4, col5, col6, col7, col8, col9}

	fmt.Println(cols)

	var (
		numBoxes int = 0
		startCol int = 0
		endCol   int = 0
	)

	var s, e = Readln(r)
	for e == nil {
		numBoxes, startCol, endCol, err = getMoveOrder(s)
		check(err)

		cols[startCol-1], cols[endCol-1] = moveBoxes3(numBoxes, cols[startCol-1], cols[endCol-1])

		//fmt.Println("Move ", numBoxes, " from ", startCol, " to ", endCol)
		s, e = Readln(r)
	}

	fmt.Println(cols)
	f.Close()
}
