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
	TotalRounds       int
	BattleListLength  int
}

type RoundRobinState struct {
	// List of items that made it to the round robin
	BattleList []*Item
	// List of all rounds to be played.
	// Slice of two Items.
	FightList [][]*Item
	// Current round robin round.
	// Used to send correct pairing to frontend.
	Current int
}
type Item struct {
	Name  string
	Score int
	// How many rounds the item has participated in.
	Rounds int
}

type Ranking struct {
	// Final list to be sent to frontend.
	// Consists of names of list items.
	RankingsList []string
	// Holder for Items when eliminated from game.
	RankingsHolder []*Item
}

// Stores the items that played in the previous round.
// Ensures there isn't an exact repeat of the previous round.
type PreviousBattlers struct {
	Battler1 *Item
	Battler2 *Item
}

// Sets the list for the current game to the user selection
func (l *BattleState) SetGame(list []*Item) {
	l.BattleList = list
	listLength := len(list)
	l.TotalRounds = 0
	l.RoundsThreshold = int(math.Floor(math.Log2(float64(listLength))) + 1)
	l.DynamicThreshold()
	RoundRobinMode = false
	l.BattleListLength = listLength
}

func (l *BattleState) SetCurrentFighters(fighters []*Item) {
	l.CurrentCombatants = fighters
}

func (l *BattleState) SetCurrentIndexes(indexes []int) {
	l.CurrentIndexes = indexes
}

// If round loser has met threshold,
// they are removed from pool of potential battlers.
func (l *BattleState) RemoveLoser(i *Item, index int) {
	l.BattleList = slices.Delete(l.BattleList, index, index+1)
	FinalRanking.AddHolder(i)
}

// Struct for receiving messages from the frontend
type Choice struct {
	Selection string `json:"selection"`
}

// Increases round winner's score and round tally.
func (i *Item) Win() {
	i.Score += 2
	i.Rounds++
}

// Checks to see if item has met threshold for removal.
func (i *Item) CheckRemoval(b *BattleState) bool {
	// Removes item if it's score is low enough,
	// Regardless of number of rounds played.
	if i.Score < -3 {
		return true
	}
	// Checks if item has played enough rounds
	// and has a low enough score to be removed from the game.
	if i.Rounds >= b.RoundsThreshold && i.Score <= b.ScoreThreshold {
		return true
	}
	return false
}

// Decreases round loser's score and increased round tally.
// Then checks if they should be removed from the pool.
func (i *Item) Lose(index int) {
	i.Score--
	i.Rounds++

	if i.CheckRemoval(&BattleList) {
		BattleList.RemoveLoser(i, index)
	}
}

// Stores eliminated item to be properly ranked at end of game.
func (r *Ranking) AddHolder(battler *Item) {
	r.RankingsHolder = append(r.RankingsHolder, battler)
}

func (r *Ranking) Final(battler string) {
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

func (b *BattleState) DynamicThreshold() {
	base := -1 * int(math.Floor(float64(len(b.BattleList))/30))
	increase := b.TotalRounds / 20 // adjust this
	b.ScoreThreshold = base + increase
}
