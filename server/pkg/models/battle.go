package models

import (
	"math/rand"
)

var BattleList = ListState{}
var FinalRanking = Ranking{}

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
		battlers[1].Lose(indexes[1])

	} else {
		battlers[1].Win()
		battlers[0].Lose(indexes[0])
	}
}
