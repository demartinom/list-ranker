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

// Begins new battle round.
func BeginRound(list []*Item) []string {
	roundConstraintCutoff := int(math.Max(6, float64(BattleList.BattleListLength)*0.33))

	// Check to see if round robin has begun.
	// If so, it begins the final round robin.
	if RoundRobinMode && RoundRobin.Current == len(RoundRobin.FightList) {
		return endGame(RoundRobin)
	}

	var (
		fighterOneIndex, fighterTwoIndex int
		fighterOne, fighterTwo           *Item
	)

	for {
		// Randomly select two items from list
		fighterOneIndex = rand.Intn(len(list))
		fighterTwoIndex = rand.Intn(len(list))

		// Ensures they are not the same item
		if fighterOneIndex == fighterTwoIndex {
			continue
		}

		fighterOne = list[fighterOneIndex]
		fighterTwo = list[fighterTwoIndex]
		// Ensures that the round is not an exact repeat of the previous round.
		if PreviousRound.Battler1 != nil && PreviousRound.Battler2 != nil {
			if fighterOne == PreviousRound.Battler1 || fighterTwo == PreviousRound.Battler2 {
				continue
			}
		}
		// If there are enough items still in the game,
		// ensures that items have played a similar number of rounds.
		if len(BattleList.BattleList) > roundConstraintCutoff && math.Abs(float64(fighterOne.Rounds)-float64(fighterTwo.Rounds)) > 2 {
			continue
		}

		break
	}
	// Set battlers about to be sent to frontend as previous battlers.
	// Used for the next round to ensure there isn't an exact repeat.
	PreviousRound = PreviousBattlers{Battler1: fighterOne, Battler2: fighterTwo}

	BattleList.SetCurrentFighters([]*Item{fighterOne, fighterTwo})
	BattleList.SetCurrentIndexes([]int{fighterOneIndex, fighterTwoIndex})

	return nil
}

// Initiated after winning item has been selected.
// Updates item scores and eliminates lose if they meet score threshold.
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

// Begins the round robin tournament.
func RoundRobinRounds(list []*Item) {
	RoundRobin.Init(list)
	RoundRobin.Current = 0
}

func endGame(rr RoundRobinState) []string {
	FinalRanking.RankingsList = nil

	sort.SliceStable(FinalRanking.RankingsHolder, func(i, j int) bool {
		s1 := float64(FinalRanking.RankingsHolder[i].Score) / float64(FinalRanking.RankingsHolder[i].Rounds)
		s2 := float64(FinalRanking.RankingsHolder[j].Score) / float64(FinalRanking.RankingsHolder[j].Rounds)
		return s1 < s2
	})

	existing := make(map[string]bool)
	for _, v := range FinalRanking.RankingsHolder {
		FinalRanking.RankingsList = append(FinalRanking.RankingsList, v.Name)
		existing[v.Name] = true
	}
	// Use SliceStable in case there are items with the same score.
	// Item eliminated first is placed lower.
	sort.SliceStable(rr.BattleList, func(i, j int) bool {
		return rr.BattleList[i].Score < rr.BattleList[j].Score
	})

	for _, v := range rr.BattleList {
		if !existing[v.Name] {
			FinalRanking.RankingsList = append(FinalRanking.RankingsList, v.Name)
		}
	}

	slices.Reverse(FinalRanking.RankingsList)
	return FinalRanking.RankingsList
}
