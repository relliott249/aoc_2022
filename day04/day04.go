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

func checkOverlap(arr [4]int) bool {
	if ((arr[0] <= arr[2]) && (arr[1] >= arr[3])) || ((arr[0]) >= arr[2] && (arr[1] <= arr[3])) {
		return true
	}
	return false
}

func checkSectOverlap(arr [4]int) bool {

	var (
		min1 int = arr[0]
		max1 int = arr[1]
		min2 int = arr[2]
		max2 int = arr[3]
	)
	if (min1 >= min2 && min1 <= max2) ||
		(max1 >= min2 && max1 <= max2) ||
		(min2 >= min1 && min2 <= max1) ||
		(max2 >= min1 && max2 <= max1) {
		return true
	}
	return false
}

func getVal(s string) (sect [4]int, err error) {
	var (
		i    int = 0
		j    int = 0
		k    int = 0
		curr     = s[i]
	)

	for curr != '-' {
		i++
		curr = s[i]
	}
	k = i + 1
	// first number
	sect[j], err = strconv.Atoi(s[:i])
	j++

	for curr != ',' {
		i++
		curr = s[i]
	}

	// second number
	sect[j], err = strconv.Atoi(s[k:i])
	k = i + 1
	j++

	for curr != '-' {
		i++
		curr = s[i]
	}
	// third number
	sect[j], err = strconv.Atoi(s[k:i])
	k = i + 1
	j++
	// fourth number
	sect[j], err = strconv.Atoi(s[k:])

	return sect, err
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

	f, err := os.Open("/home/relliott/aoc_2022/day04/input_day04.txt")
	check(err)
	r := bufio.NewReader(f)

	var s, e = Readln(r)
	var sect [4]int
	var numOverlap int = 0
	var sectOverlap int = 0

	for e == nil {
		fmt.Println(s)
		sect, err = getVal(s)
		//fmt.Printf("%+d\n", sect)
		if checkOverlap(sect) {
			numOverlap++
			//fmt.Println("overlaps!")
		}
		if checkSectOverlap(sect) {
			fmt.Println("Assignments overlap!")
			sectOverlap++
		}
		s, e = Readln(r)
	}

	fmt.Println("Number of overlapping sections is: ", numOverlap)
	fmt.Println("Number of overlapping assignments is: ", sectOverlap)

	f.Close()
}
