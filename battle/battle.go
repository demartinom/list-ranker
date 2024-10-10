package battle

type Item struct {
	Name  string
	Score int
}

func (item *Item) Win() {
	item.Score++
}

func (item *Item) Lose() {
	item.Score--
}

func Battle([]Item) {

}
