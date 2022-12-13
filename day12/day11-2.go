package main

import (
	"fmt"
)

type Monkey struct {
	Id             int
	Items          []int64
	DivBy          int64
	Children       []*Monkey
	NumInspections int
}

func (monkey *Monkey) throwItem(worryValue int64) {
	child := 0
	//worryValue = worryValue % (2*monkey.DivBy)
	if worryValue%monkey.DivBy != 0 {
		child = 1
		//fmt.Print("Worry level is not divisible by: ", monkey.DivBy, " so it throws it to ")
	}
	//fmt.Println(" monkey: ", monkey.Children[child].Id)
	monkey.Children[child].Items = append(monkey.Children[child].Items, worryValue)
	monkey.Items = monkey.Items[1:]
	monkey.NumInspections++
}

func getOp(id int, curItem int64) (worry int64) {
	switch id {
	case 0:
		worry = curItem * 19
	case 1:
		worry = curItem + 6
	case 2:
		worry = curItem * curItem
	case 3:
		worry = curItem + 3
	default:
		fmt.Errorf("Error: Invalid ID")
	}
	return worry
}

func newMonkey(id int, items []int64, divBy int64) *Monkey {
	return &Monkey{
		Id:             id,
		Items:          items,
		DivBy:          divBy,
		Children:       nil,
		NumInspections: 0,
	}
}

func (monkey *Monkey) addChild(child *Monkey) {
	monkey.Children = append(monkey.Children, child)
}

func main() {
	// initialize the monkeys
	items0 := []int64{79, 98}
	items1 := []int64{54, 65, 75, 74}
	items2 := []int64{79, 60, 97}
	items3 := []int64{74}

	monkey0 := newMonkey(0, items0, 23)
	monkey1 := newMonkey(1, items1, 19)
	monkey2 := newMonkey(2, items2, 13)
	monkey3 := newMonkey(3, items3, 17)

	monkey0.addChild(monkey2)
	monkey0.addChild(monkey3)

	monkey1.addChild(monkey2)
	monkey1.addChild(monkey0)

	monkey2.addChild(monkey1)
	monkey2.addChild(monkey3)

	monkey3.addChild(monkey0)
	monkey3.addChild(monkey1)

	monkeys := []*Monkey{monkey0, monkey1, monkey2, monkey3}

	var worryVal int64
	for round := 0; round < 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			for len(monkeys[i].Items) != 0 {
				//fmt.Println("Monkey: ", i, " inspects an item with a worry level of: ", monkeys[i].Items[0])
				worryVal = getOp(i, monkeys[i].Items[0])
				worryVal = worryVal%(23*19*13*17)
				//fmt.Println("Worry level is multiplied to be:", worryVal)
				monkeys[i].throwItem(worryVal)
			}
		}
	}
	for i := range monkeys {
		fmt.Println("Monkey: ", i, "inspected: ", monkeys[i].NumInspections, " items and has the following items: ", monkeys[i].Items)
	}

}

// guessed 6474 too low
// 10504 too low
