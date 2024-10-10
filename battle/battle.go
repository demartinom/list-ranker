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

func (item *Item) Lose() {
	item.Score--
}

func Battle(list *[]Item) {
	battlers := chooseBattlers(*list)
	var selection string

	fmt.Println("Choose which item you prefer:")
	fmt.Println("1. " + battlers[0].Name)
	fmt.Println("2. " + battlers[1].Name)

	fmt.Scanln(&selection)

	switch selection {
	case "1":
		battlers[0].Win()
		battlers[1].Lose()
	case "2":
		battlers[1].Win()
		battlers[0].Lose()
	}
}

func chooseBattlers(list []Item) []Item {
	fighterOneIndex := rand.Intn(len(list))
	fighterTwoIndex := rand.Intn(len(list))

	for fighterOneIndex == fighterTwoIndex {
		fighterTwoIndex = rand.Intn(len(list))
	}

	fighterOne := list[fighterOneIndex]
	fighterTwo := list[fighterTwoIndex]

	combatants := []Item{fighterOne, fighterTwo}
	return combatants
}
