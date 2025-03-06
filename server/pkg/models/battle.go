package models

import (
	"fmt"
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

func BattleResult(list []*Item, battlers []*Item, indexes []int, winner string) {
	if winner == battlers[0].Name {
		battlers[0].Win()
		battlers[1].Lose()
		fmt.Println(battlers[0], battlers[1])

	} else {
		battlers[1].Win()
		battlers[0].Lose()
		fmt.Println(battlers[0], battlers[1])
	}
}
