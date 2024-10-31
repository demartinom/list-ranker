package battle

func Battle(list *[]Item, input input, output output) {
	var results []string
	for len(*list) > 1 {
		battlers, indexes := input.chooseBattlers(*list)
		output.RemainingItems(list)
		output.Fight(battlers, indexes, list, &results)
	}
	output.Result(results, list)
}
