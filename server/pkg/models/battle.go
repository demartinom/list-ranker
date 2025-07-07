package models

import (
	"math"
	"math/rand"
	"slices"
	"sort"
)

var (
	BattleList     = BattleState{}
	FinalRanking   = Ranking{}
	PreviousRound  = PreviousBattlers{}
	RoundRobin     = RoundRobinState{}
	RoundRobinMode = false
)

func BeginRound(list []*Item) []string {
	roundConstraintCutoff := int(float64(BattleList.BattleListLength) * 0.33)
	if RoundRobinMode && RoundRobin.Current == len(RoundRobin.FightList) {
		return endGame(RoundRobin)
	}

	var (
		fighterOneIndex, fighterTwoIndex int
		fighterOne, fighterTwo           *Item
	)

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

		if len(BattleList.BattleList) > roundConstraintCutoff && math.Abs(float64(fighterOne.Rounds)-float64(fighterTwo.Rounds)) > 2 {
			continue
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
	BattleList.TotalRounds++
	BattleList.DynamicThreshold()
}

func RoundRobinRounds(list []*Item) {
	RoundRobin.Init(list)
	RoundRobin.Current = 0
}

func endGame(rr RoundRobinState) []string {
	sort.Slice(rr.BattleList, func(i, j int) bool {
		return rr.BattleList[i].Score < rr.BattleList[j].Score
	})
	for _, v := range rr.BattleList {
		FinalRanking.AddItem(v.Name)
	}
	slices.Reverse(FinalRanking.RankingsList)
	return FinalRanking.RankingsList
}
