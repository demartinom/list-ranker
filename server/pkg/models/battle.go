package models

import (
	"math/rand"
	"slices"
)

var BattleList = ListState{}
var FinalRanking = Ranking{}

func BeginRound(list []*Item) []string {
	if len(list) == 1 {
		FinalRanking.AddItem(list[0].Name)
		final := endGame()
		return final
	}

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

	return nil
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

func endGame() []string {
	slices.Reverse(FinalRanking.RankingsList)
	return FinalRanking.RankingsList
}
