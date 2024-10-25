package battle

type input interface {
	chooseBattlers(list []Item) ([]*Item, []int)
}
type Item struct {
	Name  string
	Score int
}

// Increment winner of battles score. Capped at 5 to increase game speed.
func (item *Item) Win() {
	if item.Score < 5 {
		item.Score++
	}
}

// Decrement loser of battles score.
// If score drops below -2 they will be removed from the game
func (item *Item) Lose(list *[]Item, index int, results *[]string) {
	item.Score--
	if len(*list) == 2 {
		RemoveLoser(list, index, results)
	}
	if item.Score <= -2 {
		RemoveLoser(list, index, results)
	}
}
