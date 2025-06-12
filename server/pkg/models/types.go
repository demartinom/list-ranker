package models

// Create global state for list selected for battle
type BattleState struct {
	BattleList        []*Item
	CurrentCombatants []*Item
	CurrentIndexes    []int
	RoundsThreshold   int
	ScoreThreshold    int
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
func (l *BattleState) SetList(list []*Item) {
	l.BattleList = list
}

func (l *BattleState) SetCurrentFighters(fighters []*Item) {
	l.CurrentCombatants = fighters
}

func (l *BattleState) SetCurrentIndexes(indexes []int) {
	l.CurrentIndexes = indexes
}

func (l *BattleState) RemoveLoser(i *Item, index int) {
	l.BattleList = append(l.BattleList[:index], l.BattleList[index+1:]...)
	FinalRanking.AddItem(i.Name)
}

// Struct for receiving messages from the frontend
type Choice struct {
	Selection string `json:"selection"`
}

func (i *Item) Win() {
	i.Score++
}

func (i *Item) Lose(index int) {
	i.Score--
	if len(BattleList.BattleList) == 2 {
		BattleList.RemoveLoser(i, index)
	}
	if i.Score <= -2 {
		BattleList.RemoveLoser(i, index)
	}
}

func (r *Ranking) AddItem(battler string) {
	r.RankingsList = append(r.RankingsList, battler)
}
