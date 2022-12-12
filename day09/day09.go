package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pos struct {
	x, y int
}

func main() {
	//Read input file
	input, _ := os.Open("/home/relliott/aoc_2022/day09/input_day09.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	// fist time looking into map structure
	visitedByTail := make(map[pos]bool)
	knots := make([]pos, 10)

	//head := pos{0, 0}
	//tail := pos{0, 0}
	//visitedByTail[tail] = true
	visitedByTail[knots[9]] = true

	for sc.Scan() {
		direction := rune(sc.Text()[0])
		moves, _ := strconv.Atoi(sc.Text()[2:])

		//I calculate moves one by one
		for moves > 0 {
			switch direction {
			case 'U':
				knots[0].y++
			case 'R':
				knots[0].x++
			case 'D':
				knots[0].y--
			case 'L':
				knots[0].x--
			}
			for i := range knots[:len(knots)-1] {
				knots[i+1] = moveTails(knots[i+1], knots[i])
			}
			moves--
			//tail = moveTail(tail, head)
			//visitedByTail[tail] = true
			visitedByTail[knots[9]] = true
		}
	}

	fmt.Println(len(visitedByTail))
}

func moveTail(tail pos, head pos) (newTail pos) {
	newTail = tail
	switch (pos{head.x - tail.x, head.y - tail.y}) {
	case pos{-2, 1}, pos{-1, 2}, pos{0, 2}, pos{1, 2}, pos{2, 1}:
		newTail.y++
	}
	switch (pos{head.x - tail.x, head.y - tail.y}) {
	case pos{1, 2}, pos{2, 1}, pos{2, 0}, pos{2, -1}, pos{1, -2}:
		newTail.x++
	}
	switch (pos{head.x - tail.x, head.y - tail.y}) {
	case pos{2, -1}, pos{1, -2}, pos{0, -2}, pos{-1, -2}, pos{-2, -1}:
		newTail.y--
	}
	switch (pos{head.x - tail.x, head.y - tail.y}) {
	case pos{-1, -2}, pos{-2, -1}, pos{-2, -0}, pos{-2, 1}, pos{-1, 2}:
		newTail.x--
	}
	return
}

func moveTails(tail pos, head pos) (newTail pos) {
	newTail = tail
	switch (pos{head.x - tail.x, head.y - tail.y}) {
	case pos{-2, 1}, pos{-1, 2}, pos{0, 2}, pos{1, 2}, pos{2, 1}, pos{2, 2}, pos{-2, 2}:
		newTail.y++
	}
	switch (pos{head.x - tail.x, head.y - tail.y}) {
	case pos{1, 2}, pos{2, 1}, pos{2, 0}, pos{2, -1}, pos{1, -2}, pos{2, 2}, pos{2, -2}:
		newTail.x++
	}
	switch (pos{head.x - tail.x, head.y - tail.y}) {
	case pos{-2, -2}, pos{2, -1}, pos{1, -2}, pos{0, -2}, pos{-1, -2}, pos{-2, -1}, pos{2, -2}:
		newTail.y--
	}
	switch (pos{head.x - tail.x, head.y - tail.y}) {
	case pos{-2, -2}, pos{-1, -2}, pos{-2, -1}, pos{-2, -0}, pos{-2, 1}, pos{-1, 2}, pos{-2, 2}:
		newTail.x--
	}
	return
}
