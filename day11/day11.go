package main

import "fmt"

type Monkey struct {
	Id             int
	Items          []int
	DivBy          int
	Children       []*Monkey
	NumInspections int
}

func (monkey *Monkey) throwItem(worryValue int) {
	child := 0
	if worryValue%monkey.DivBy != 0 {
		child = 1
	}
	monkey.Children[child].Items = append(monkey.Children[child].Items, worryValue)
	monkey.Items = monkey.Items[1:]
	monkey.NumInspections++
}

func getOp(id int, curItem int) (worry int) {
	switch id {
	case 0:
		worry = curItem * 19
	case 1:
		worry = curItem + 1
	case 2:
		worry = curItem + 6
	case 3:
		worry = curItem + 5
	case 4:
		worry = curItem * curItem
	case 5:
		worry = curItem + 7
	case 6:
		worry = curItem * 7
	case 7:
		worry = curItem + 2
	default:
		fmt.Errorf("Error: Invalid ID")
	}
	return worry
}

func newMonkey(id int, items []int, divBy int) *Monkey {
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
	items0 := []int{57, 58}
	items1 := []int{66, 52, 59, 79, 94, 73}
	items2 := []int{80}
	items3 := []int{82, 81, 68, 66, 71, 83, 75, 97}
	items4 := []int{55, 52, 67, 70, 69, 94, 90}
	items5 := []int{69, 85, 89, 91}
	items6 := []int{75, 53, 73, 52, 75}
	items7 := []int{94, 60, 79}

	monkey0 := newMonkey(0, items0, 7)
	monkey1 := newMonkey(1, items1, 19)
	monkey2 := newMonkey(2, items2, 5)
	monkey3 := newMonkey(3, items3, 11)
	monkey4 := newMonkey(4, items4, 17)
	monkey5 := newMonkey(5, items5, 13)
	monkey6 := newMonkey(6, items6, 2)
	monkey7 := newMonkey(7, items7, 3)

	monkey0.addChild(monkey2)
	monkey0.addChild(monkey3)

	monkey1.addChild(monkey4)
	monkey1.addChild(monkey6)

	monkey2.addChild(monkey7)
	monkey2.addChild(monkey5)

	monkey3.addChild(monkey5)
	monkey3.addChild(monkey2)

	monkey4.addChild(monkey0)
	monkey4.addChild(monkey3)

	monkey5.addChild(monkey1)
	monkey5.addChild(monkey7)

	monkey6.addChild(monkey0)
	monkey6.addChild(monkey4)

	monkey7.addChild(monkey1)
	monkey7.addChild(monkey6)

	monkeys := []*Monkey{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}
	totalFactor := 1
	for i:= 0; i < len(monkeys); i++{
		totalFactor *= monkeys[i].DivBy
	}

	var worryVal int
	for round := 0; round < 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			for len(monkeys[i].Items) != 0 {
				worryVal = getOp(i, monkeys[i].Items[0]) % totalFactor
				monkeys[i].throwItem(worryVal)
			}
		}
	}
	for i := range monkeys {
		fmt.Println("Monkey: ", i, "inspected: ", monkeys[i].NumInspections, " items and has the following items: ", monkeys[i].Items)
	}

}

// guessed 6474 too low
