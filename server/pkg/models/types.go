package models

// Create global state for list selected for battle
type ListState struct {
	BattleList        []*Item
	CurrentCombatants []*Item
	CurrentIndexes    []int
}

type Item struct {
	Name  string
	Score int
}

type Ranking struct {
	RankingsList []string
}

// Sets the list for the current game to the user selection
func (l *ListState) SetList(list []*Item) {
	l.BattleList = list
}

func (l *ListState) SetCurrentFighters(fighters []*Item) {
	l.CurrentCombatants = fighters
}

func (l *ListState) SetCurrentIndexes(indexes []int) {
	l.CurrentIndexes = indexes
}

func (l *ListState) RemoveLoser(i *Item, index int) {
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
