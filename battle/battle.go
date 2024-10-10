package battle

import "math/rand"

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

func Battle(list []Item) {
	battlers := chooseBattlers(list)
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
