package filehandler

import "github.com/demartinom/list-ranker/battle"

func ConvertToSlice(listInput [][]string) []battle.Item {
	var itemsList []battle.Item

	for _, itemInput := range listInput {
		itemsList = append(itemsList, battle.Item{Name: itemInput[0], Score: 0})
	}

	return itemsList
}
