package filehandler

type Item struct {
	Name  string
	Score int
}

func ConvertToSlice(listInput [][]string) []Item {
	var itemsList []Item

	for _, itemInput := range listInput {
		itemsList = append(itemsList, Item{itemInput[0], 0})
	}

	return itemsList
}
