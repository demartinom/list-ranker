package battle

import "fmt"

func Battle(list *[]Item, input input, output output) {
	var results []string
	for len(*list) > 1 {
		battlers, indexes := input.chooseBattlers(*list)
		output.RemainingItems(list)
		output.Fight(battlers, indexes, list, &results)
	}
	output.Result(results, list)
}

func RemoveLoser(list *[]Item, index int, results *[]string) {
	placement := fmt.Sprintf("%d: %s", len(*list), (*list)[index].Name)
	*results = append(*results, placement)
	*list = append((*list)[:index], (*list)[index+1:]...)
}
