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

type Directory struct {
	name     string
	parent   *Directory
	contents []string
	child    *Directory
	size     int
}

func main() {

	f, err := os.Open("/home/relliott/aoc_2022/day07/input_day07.txt")
	check(err)
	r := bufio.NewReader(f)

	var i int = 0
	var dirList [10]Directory
	var s, e = Readln(r)
	for e == nil {
		currLine := strings.Fields(s)
		var fileSize int
		if currLine[0] == "$" {
			if currLine[1] == "cd" {
				if currLine[2] == ".." {
					i--
				} else {
					dirList[i].Name = currLine[2]
					i++
				}
			} else if currLine[1] == "ls" {
				s, e = Readln(r)
				currLine = strings.Fields(s)
				for currLine[1] != "$" {
					dirList[i].contents = append(dirList[i].contents, currLine[1])
					fileSize, err = strconv.Atoi(strings.TrimSpace(currLine[1]))
					dirList[i].size += fileSize
					s, e = Readln(r)
					currLine = strings.Fields(s)
				}
			}
		}
		s, e = Readln(r)
	}
	fmt.Println(dirList)
}
