package models

// Create global state for list selected for battle
type ListState struct {
	BattleList        []*Item
	CurrentCombatants []*Item
	CurrentIndexes    []int
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

// Struct for receiving messages from the frontend
type Choice struct {
	Selection string `json:"selection"`
}

type Item struct {
	Name  string
	Score int
}

func (i *Item) Win() {
	i.Score++
}

func (i *Item) Lose() {
	i.Score--
}
