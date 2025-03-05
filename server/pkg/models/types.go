package models

type ListState struct {
	BattleList []*Item
}

func (l *ListState) SetList(list []*Item) {
	l.BattleList = list
}

type Choice struct {
	Selection string `json:"selection"`
}

type Item struct {
	Name  string
	Score int
}
