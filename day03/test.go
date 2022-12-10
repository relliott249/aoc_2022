package main

import (
	"fmt"
)

func main() {
	var secretString string = "this is a top secret string"

	res := secretString[0:10]

	res2 := secretString[:5]

	res3 := secretString[10:]

	fmt.Println(res, res2, res3)

	half := secretString[:(len(secretString) / 2)]

	fmt.Println(half)

	for i := 0; i < len(half); i++ {
		for x := len(half); x < len(secretString); x++ {
			if half[i] == secretString[x] {
				fmt.Printf("%c found in %s\n", half[i], secretString)
			}
		}
	}
}
