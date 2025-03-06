package models

import (
	"math/rand"
)

var BattleList = ListState{}

func ChooseBattlers(list []*Item) {
	fighterOneIndex := rand.Intn(len(list))
	fighterTwoIndex := rand.Intn(len(list))

	for fighterOneIndex == fighterTwoIndex {
		fighterTwoIndex = rand.Intn(len(list))
	}

	fighterOne := list[fighterOneIndex]
	fighterTwo := list[fighterTwoIndex]

	combatants := []*Item{fighterOne, fighterTwo}
	indexes := []int{fighterOneIndex, fighterTwoIndex}

	BattleList.SetCurrentFighters(combatants)
	BattleList.SetCurrentIndexes(indexes)
}
