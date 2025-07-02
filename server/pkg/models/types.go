package models

import (
	"math"
	"slices"
)

// Create global state for list selected for battle
type BattleState struct {
	BattleList        []*Item
	CurrentCombatants []*Item
	CurrentIndexes    []int
	RoundsThreshold   int
	ScoreThreshold    int
	StaleElimination  int
}

type RoundRobinState struct {
	BattleList []*Item
	FightList  [][]*Item
	Current    int
}
type Item struct {
	Name   string
	Score  int
	Rounds int
}

type Ranking struct {
	RankingsList []string
}

type PreviousBattlers struct {
	Battler1 *Item
	Battler2 *Item
}

// Sets the list for the current game to the user selection
func (l *BattleState) SetGame(list []*Item) {
	l.BattleList = list
	listLength := len(list)

	l.RoundsThreshold = int(math.Floor(math.Log2(float64(listLength))) + 1)
	l.ScoreThreshold = -1 * int(math.Max(2, math.Floor(float64(listLength)/20)))
	RoundRobinMode = false
	l.StaleElimination = max(int(math.Round(math.Log2(float64(listLength)))), 1)
}

func (l *BattleState) SetCurrentFighters(fighters []*Item) {
	l.CurrentCombatants = fighters
}

func (l *BattleState) SetCurrentIndexes(indexes []int) {
	l.CurrentIndexes = indexes
}

func (l *BattleState) RemoveLoser(i *Item, index int) {
	l.BattleList = slices.Delete(l.BattleList, index, index+1)
	FinalRanking.AddItem(i.Name)
}

// Struct for receiving messages from the frontend
type Choice struct {
	Selection string `json:"selection"`
}

func (i *Item) Win() {
	i.Score += 2
	i.Rounds++
}

func (i *Item) CheckRemoval(b *BattleState) bool {
	if i.Score < -5 {
		return true
	}

	if i.Rounds > 8 && i.Score < BattleList.StaleElimination {
		return true
	} else if i.Rounds >= b.RoundsThreshold && i.Score <= b.ScoreThreshold {
		return true
	}
	return false
}

func (i *Item) Lose(index int) {
	i.Score--
	i.Rounds++

	if i.CheckRemoval(&BattleList) {
		BattleList.RemoveLoser(i, index)
	}
}

func (r *Ranking) AddItem(battler string) {
	r.RankingsList = append(r.RankingsList, battler)
}

func (rr *RoundRobinState) Init(list []*Item) {
	var pairings [][]*Item

	for i := 0; i < len(list); i++ {
		p1 := list[i]
		p1.Score = 0
		for j := i + 1; j < len(list); j++ {
			p2 := list[j]
			p2.Score = 0
			pairings = append(pairings, []*Item{p1, p2})
		}
	}
	rr.FightList = pairings
	rr.BattleList = list
}

func (rr *RoundRobinState) RRRound(winner string) {
	fighter1 := RoundRobin.FightList[RoundRobin.Current][0]
	fighter2 := RoundRobin.FightList[RoundRobin.Current][1]

	if winner == fighter1.Name {
		fighter1.Score++
	} else {
		fighter2.Score++
	}
	rr.Current++
}
