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

func checkScore(s string) int {
	var (
		score        int  = 0
		win          int  = 6
		draw         int  = 3
		lose         int  = 0
		elfRock      byte = 'A'
		elfPaper     byte = 'B'
		elfScissors  byte = 'C'
		selfRock     byte = 'X'
		selfPaper    byte = 'Y'
		selfScissors byte = 'Z'
		rock         int  = 1
		paper        int  = 2
		scissors     int  = 3
	)
	if len(s) != 3 {
		return 0
	}

	var elf = s[0]
	var me = s[2]

	// Case of a draw
	if me == selfRock {
		if elf == elfPaper {
			fmt.Println("Lose %c: Paper > Rock ")
			score = lose + rock
		} else if elf == elfScissors {
			fmt.Println("Win C X: Rock > Scissors")
			score = win + rock
		} else {
			fmt.Println("Draw")
			score = draw + rock
		}
	} else if me == selfPaper {
		if elf == elfScissors {
			fmt.Println("Lose C Y: Scissors > Paper")
			score = lose + paper
		} else if elf == elfRock {
			fmt.Println("Win A Y: Rock > Paper")
			score = win + paper
		} else {
			fmt.Println("Draw")
			score = draw + paper
		}
		// self scissors
	} else if me == selfScissors {
		if elf == elfRock {
			fmt.Println("Lose A Z: Rock > Scissors")
			score = lose + scissors
		} else if elf == elfPaper {
			fmt.Println("Win B Z: Scissors > paper")
			score = win + scissors
		} else {
			fmt.Println("Draw")
			score = draw + scissors
		}
	}
	return score
}

func partTwo(s string) int {
	var (
		rock     byte = 'A'
		paper    byte = 'B'
		scissors byte = 'C'
		lose     byte = 'X'
		//draw          byte = 'Y'
		win           byte = 'Z'
		rockValue     int  = 1
		paperValue    int  = 2
		scissorsValue int  = 3
		winValue      int  = 6
		loseValue     int  = 0
		drawValue     int  = 3
		score         int  = 0
	)
	if len(s) != 3 {
		return 0
	}

	var elf = s[0]
	var outcome = s[2]

	if elf == rock {
		if outcome == win {
			score = winValue + paperValue
		} else if outcome == lose {
			score = loseValue + scissorsValue
		} else {
			score = drawValue + rockValue
		}
	} else if elf == paper {
		if outcome == win {
			score = winValue + scissorsValue
		} else if outcome == lose {
			score = loseValue + rockValue
		} else {
			score = drawValue + paperValue
		}
	} else if elf == scissors {
		if outcome == win {
			score = winValue + rockValue
		} else if outcome == lose {
			score = loseValue + paperValue
		} else {
			score = drawValue + scissorsValue
		}
	}
	return score
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

	var (
		currentScore int = 0
	)

	f, err := os.Open("/home/relliott/aoc_2022/day02/input_day02.txt")
	check(err)

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	r := bufio.NewReader(f)
	var s, e = Readln(r)
	for e == nil {
		//fmt.Println(s)
		currentScore += partTwo(s)
		s, e = Readln(r)
	}

	fmt.Println(currentScore)

	f.Close()
}
