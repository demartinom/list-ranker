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
