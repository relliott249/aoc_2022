package main

import (
	"bufio"
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

func numVisible(s []string)(num int){
	for _, range:= len(s)
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



}
