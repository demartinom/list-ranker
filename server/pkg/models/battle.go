package models

import (
	"math/rand"
	"slices"
)

var BattleList = BattleState{}
var FinalRanking = Ranking{}
var PreviousRound = PreviousBattlers{}

func BeginRound(list []*Item) []string {
	if len(list) == 1 {
		FinalRanking.AddItem(list[0].Name)
		return endGame()
	}

	var fighterOneIndex, fighterTwoIndex int
	var fighterOne, fighterTwo *Item

	for {
		fighterOneIndex = rand.Intn(len(list))
		fighterTwoIndex = rand.Intn(len(list))

		if fighterOneIndex == fighterTwoIndex {
			continue
		}

		fighterOne = list[fighterOneIndex]
		fighterTwo = list[fighterTwoIndex]

		if PreviousRound.Battler1 != nil && PreviousRound.Battler2 != nil {
			if fighterOne == PreviousRound.Battler1 || fighterTwo == PreviousRound.Battler2 {
				continue
			}
		}

		break
	}

	PreviousRound = PreviousBattlers{Battler1: fighterOne, Battler2: fighterTwo}

	BattleList.SetCurrentFighters([]*Item{fighterOne, fighterTwo})
	BattleList.SetCurrentIndexes([]int{fighterOneIndex, fighterTwoIndex})

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
