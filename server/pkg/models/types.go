package models

// Create global state for list selected for battle
type ListState struct {
	BattleList []*Item
}

// Sets the list for the current game to the user selection
func (l *ListState) SetList(list []*Item) {
	l.BattleList = list
}

// Struct for receiving messages from the frontend
type Choice struct {
	Selection string `json:"selection"`
}

type Item struct {
	Name  string
	Score int
}
