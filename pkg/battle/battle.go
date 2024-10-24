package battle

func Battle(list *[]Item, input input, output output) {

}

// // Increment winner of battles score. Capped at 5 to increase game speed.
// func (item *Item) Win() {
// 	if item.Score < 5 {
// 		item.Score++
// 	}
// }

// // Decrement loser of battles score.
// // If score drops below -2 they will be removed from the game
// func (item *Item) Lose(list *[]Item, index int, results *[]string) {
// 	item.Score--
// 	if len(*list) == 2 {
// 		RemoveLoser(list, index, results)
// 	}
// 	if item.Score <= -2 {
// 		RemoveLoser(list, index, results)
// 	}
// }

// func Battle(list *[]Item) {
// 	var results []string

// 	for len(*list) > 1 {
// 		battlers, indexes := chooseBattlers(*list)
// 		var selection string

// 		fmt.Printf("Remaining items : %d\n", len(*list))

// 		fmt.Println("Choose which item you prefer:")
// 		fmt.Println("1. " + battlers[0].Name)
// 		fmt.Println("2. " + battlers[1].Name)

// 	selectloop:
// 		for {
// 			fmt.Scanln(&selection)

// 			switch selection {
// 			case "1":
// 				battlers[0].Win()
// 				battlers[1].Lose(list, indexes[1], &results)
// 				break selectloop
// 			case "2":
// 				battlers[1].Win()
// 				battlers[0].Lose(list, indexes[0], &results)
// 				break selectloop
// 			default:
// 				fmt.Println("Invalid input")
// 				fmt.Println("1. " + battlers[0].Name)
// 				fmt.Println("2. " + battlers[1].Name)
// 			}
// 		}
// 	}
// 	endResult(results, list)
// }

// // Selects two items at random from pool of items.
// // Ensures the same item isn't listed as both options
// func chooseBattlers(list []Item) ([]*Item, []int) {
// 	fighterOneIndex := rand.Intn(len(list))
// 	fighterTwoIndex := rand.Intn(len(list))

// 	for fighterOneIndex == fighterTwoIndex {
// 		fighterTwoIndex = rand.Intn(len(list))
// 	}

// 	fighterOne := &list[fighterOneIndex]
// 	fighterTwo := &list[fighterTwoIndex]

// 	combatants := []*Item{fighterOne, fighterTwo}
// 	indexes := []int{fighterOneIndex, fighterTwoIndex}
// 	return combatants, indexes
// }

// // When score threshold met, item is removed from pool of items
// // and added to final ranking
// func RemoveLoser(list *[]Item, index int, results *[]string) {
// 	placement := fmt.Sprintf("%d: %s", len(*list), (*list)[index].Name)
// 	*results = append(*results, placement)
// 	*list = append((*list)[:index], (*list)[index+1:]...)
// }

// // Prints out the final item rankings
// func endResult(results []string, list *[]Item) {
// 	results = append(results, fmt.Sprintf("1. %s", (*list)[0].Name))
// 	slices.Reverse(results)
// 	fmt.Println("Your Results:")
// 	for _, item := range results {
// 		fmt.Println(item)
// 	}
// }
