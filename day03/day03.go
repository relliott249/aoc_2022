package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func calcPriority(s int) int {
	var (
		upper2Value int = 38
		lower2Value int = 96
		retValue    int = 0
	)
	// in the ascii table upper case letters come before a
	if s <= 'a' {
		retValue = int(s) - upper2Value
	} else {
		retValue = int(s) - lower2Value
	}

	return retValue

}

func findMatch(s byte, s2 string) bool {
	for i := 0; i < len(s2); i++ {
		if s == s2[i] {
			return true
		}
	}
	return false
}

func findMatches(s byte, s2 string, s3 string) bool {
	for i := 0; i < len(s2); i++ {
		if s == s2[i] {
			for j := 0; j < len(s3); j++ {
				if s == s3[j] {
					return true
				}
			}
		}
	}
	return false
}

func splitString(s string) (string, string) {
	var stringLength int = len(s)
	var s1 = s[:stringLength/2]
	var s2 = s[stringLength/2:]
	return s1, s2
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

	f, err := os.Open("/home/relliott/aoc_2022/day03/input_day03.txt")
	check(err)

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	var s1 string
	var s2 string
	var charValue int
	var totalValue int = 0
	r := bufio.NewReader(f)
	var s, e = Readln(r)
	for e == nil {
		s1, s2 = splitString(s)
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
		for i := 0; i < len(s1); i++ {
			if findMatch(s1[i], s2) {
				fmt.Printf("Found priority package %c\n", s1[i])
				charValue = calcPriority(int(s1[i]))
				fmt.Println("With Value: ", charValue)
				totalValue += charValue
				break
			}
		}
		s1, e = Readln(r)
	}

	fmt.Println("Total value for priority packages is: ", totalValue)
}
