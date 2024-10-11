package battle

import (
	"fmt"
	"math/rand"
)

type Item struct {
	Name  string
	Score int
}

func (item *Item) Win() {
	item.Score++
}

func (item *Item) Lose(list *[]Item, index int) {
	item.Score--
	if item.Score <= -2 {
		RemoveLoser(list, index)
	}
}

func Battle(list *[]Item) {
	for len(*list) > 1 {
		battlers, indexes := chooseBattlers(*list)
		var selection string

		fmt.Println("Choose which item you prefer:")
		fmt.Println("1. " + battlers[0].Name)
		fmt.Println("2. " + battlers[1].Name)

	selectloop:
		for {
			fmt.Scanln(&selection)

			switch selection {
			case "1":
				battlers[0].Win()
				battlers[1].Lose(list, indexes[1])
				fmt.Println(list)
				break selectloop
			case "2":
				battlers[1].Win()
				battlers[0].Lose(list, indexes[0])
				fmt.Println(list)
				break selectloop
			default:
				fmt.Println("Invalid input")
				fmt.Println("1. " + battlers[0].Name)
				fmt.Println("2. " + battlers[1].Name)
			}
		}
	}
}

func chooseBattlers(list []Item) ([]*Item, []int) {
	fighterOneIndex := rand.Intn(len(list))
	fighterTwoIndex := rand.Intn(len(list))

	for fighterOneIndex == fighterTwoIndex {
		fighterTwoIndex = rand.Intn(len(list))
	}

	fighterOne := &list[fighterOneIndex]
	fighterTwo := &list[fighterTwoIndex]

	combatants := []*Item{fighterOne, fighterTwo}
	indexes := []int{fighterOneIndex, fighterTwoIndex}
	return combatants, indexes
}

func RemoveLoser(list *[]Item, index int) {
	*list = append((*list)[:index], (*list)[index+1:]...)
}
