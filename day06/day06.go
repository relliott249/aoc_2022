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

func checkMarker(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func findKeyPos(s string, keyLen int) (numChars int) {
	for i := 0; i < (len(s) - keyLen); i++ {
		if checkMarker(s[i : i+keyLen]) {
			numChars = i
			break
		}
	}
	return numChars
}

func main() {

	f, err := os.Open("/home/relliott/aoc_2022/day06/input_day06.txt")
	check(err)
	r := bufio.NewReader(f)
	var numChars int = 0
	var keyLen = 14
	var s, e = Readln(r)
	for e == nil {
		numChars = findKeyPos(s, keyLen) + keyLen
		s, e = Readln(r)
	}
	fmt.Println("First key appears after: ", numChars, " characters")
	f.Close()
}
