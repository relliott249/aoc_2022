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

type stateMachine struct {
	Cycle int
	RegX  int
}

func drawSprite(regVal []int) {
	var crtLine string = ""
	for i := 0; i < 40; i++ {
		if i == regVal[i] || i == regVal[i]-1 || i == regVal[i]+1 {
			crtLine += "#"
		} else {
			crtLine += "."
		}
	}
	fmt.Println(crtLine)
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
	f, err := os.Open("/home/relliott/aoc_2022/day10/input_day10.txt")
	check(err)
	r := bufio.NewReader(f)
	//var dirList [10]Directory
	var aocInput []string
	var s, e = Readln(r)
	aocInput = append(aocInput, s)
	for e == nil {
		s, e = Readln(r)
		aocInput = append(aocInput, s)
	}
	aocInput = aocInput[:len(aocInput)-1]
	f.Close()

	// initial value of register is 1
	var regX []int
	regX = append(regX, 1)
	var cycle int = 0
	for i := 0; i < len(aocInput); i++ {
		curLine := strings.Fields(aocInput[i])
		switch curLine[0] {
		case "noop":
			regX = append(regX, regX[cycle])
			cycle++
		case "addx":
			regX = append(regX, regX[cycle])
			cycle++
			addValue, e := strconv.Atoi(curLine[1])
			check(e)
			regX = append(regX, regX[cycle]+addValue)
			cycle++
		default:
			fmt.Errorf("Error: Invalid command")

		}
	}
	//	var totalStr int = 0
	//	for i := 19; i < len(regX); i += 40 {
	//		totalStr += (i + 1) * regX[i]
	//		fmt.Println("Cycle: ", i+1, " X: ", regX[i], " Strength: ", totalStr)
	//	}
	//for i := 0; i < len(regX); i++ {
	//	fmt.Println("Cycle: ", i+1, " X: ", regX[i])
	//}
	for i := 0; i < len(regX)-40; i += 40 {
		drawSprite(regX[i : i+40])
	}
}
