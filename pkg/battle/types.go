package battle

import (
	"fmt"
	"math/rand"
)

type input interface {
	chooseBattlers(list []Item) ([]*Item, []int)
}

type CLIInput struct{}

func (c CLIInput) chooseBattlers(list []Item) ([]*Item, []int) {
	fighterOneIndex := rand.Intn(len(list))
	fighterTwoIndex := rand.Intn(len(list))

	for fighterOneIndex == fighterTwoIndex {
		fighterTwoIndex = rand.Intn(len(list))
	}

	fighterOne := &list[fighterOneIndex]
	fighterTwo := &list[fighterTwoIndex]

	combatants := []*Item{fighterOne, fighterTwo}
	indexes := []int{fighterOneIndex, fighterTwoIndex}
	return combatants, indexes
}

type output interface {
	RemainingItems(count int)
}

type CLIOutput struct{}

func (c CLIOutput) RemainingItems(count int) {
	fmt.Printf("Remaining items : %d\n", count)
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
