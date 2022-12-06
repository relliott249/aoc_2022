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

// AddCalories takes the string for the current line in the file
// and the current elfs calorie count and adds them together
func AddCalories(currentCount int, s string) (int, error) {
	var (
		currentCalories       = 0
		err             error = nil
	)
	currentCalories, err = strconv.Atoi(s)
	check(err)
	currentCount = currentCalories + currentCount
	return currentCount, err
}

func main() {
	var (
		currentElf      = 0
		currentCalories = 0
		maxCalories0    = 0
		maxCalories1    = 0
		maxCalories2    = 0
	)

	f, err := os.Open("./input_day1.txt")
	check(err)

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	r := bufio.NewReader(f)
	var s, e = Readln(r)
	for e == nil {
		// if the current line isn't empty
		if 0 != len(s) {
			currentCalories, err = AddCalories(currentCalories, s)

		} else {
			if currentCalories > maxCalories0 {
				maxCalories0 = currentCalories
			} else if currentCalories > maxCalories1 {
				maxCalories1 = currentCalories
			} else if currentCalories > maxCalories2 {
				maxCalories2 = currentCalories
			}
			fmt.Println("Elf: ", currentElf, " calories: ", currentCalories)
			currentElf++
			currentCalories = 0
		}
		s, e = Readln(r)
	}
	fmt.Println("Most calories carried by an elf: ", maxCalories0)
	fmt.Println("Second most calories carried by an elf: ", maxCalories1)
	fmt.Println("Third most calories carried by an elf: ", maxCalories2)
	var total_calories = maxCalories0 + maxCalories1 + maxCalories2
	fmt.Println("For a total of: ", total_calories)

	f.Close()
}
