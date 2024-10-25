package battle

import "fmt"

func ConvertToSlice(listInput [][]string) []Item {
	var itemsList []Item

	for _, itemInput := range listInput {
		itemsList = append(itemsList, Item{Name: itemInput[0], Score: 0})
	}

	return itemsList
}

func RemoveLoser(list *[]Item, index int, results *[]string) {
	placement := fmt.Sprintf("%d: %s", len(*list), (*list)[index].Name)
	*results = append(*results, placement)
	*list = append((*list)[:index], (*list)[index+1:]...)
}
