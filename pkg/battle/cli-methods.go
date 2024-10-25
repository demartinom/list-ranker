package battle

import (
	"fmt"
	"math/rand"
	"slices"
)

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
	RemainingItems(list *[]Item)
	Fight(battlers []*Item, indexes []int, list *[]Item, results *[]string)
	Result(results []string, list *[]Item)
}

type CLIOutput struct{}

func (c CLIOutput) Fight(battlers []*Item, indexes []int, list *[]Item, results *[]string) {
	var selection string

	fmt.Println("Choose which item you prefer:")
	fmt.Println("1. " + battlers[0].Name)
	fmt.Println("2. " + battlers[1].Name)

	for {
		fmt.Scanln(&selection)

		switch selection {
		case "1":
			battlers[0].Win()
			battlers[1].Lose(list, indexes[1], results)
			return
		case "2":
			battlers[1].Win()
			battlers[0].Lose(list, indexes[0], results)
			return
		default:
			fmt.Println("Invalid input")
			fmt.Println("1. " + battlers[0].Name)
			fmt.Println("2. " + battlers[1].Name)
		}
	}

}

func (c CLIOutput) RemainingItems(list *[]Item) {
	fmt.Printf("Remaining items : %d\n", len(*list))
}

func (c CLIOutput) Result(results []string, list *[]Item) {
	results = append(results, fmt.Sprintf("1. %s", (*list)[0].Name))
	slices.Reverse(results)
	fmt.Println("Your Results:")
	for _, item := range results {
		fmt.Println(item)
	}
}
