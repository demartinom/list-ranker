package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func ReadCSV(fileName string) []*Item {
	filePath := fmt.Sprintf("game-data/%s.csv", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	//Skip over header line
	reader.Read()

	listItems, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []*Item

	for _, itemInput := range listItems {
		itemsList = append(itemsList, &Item{Name: itemInput[0], Score: 0})
	}

	return itemsList
}

func ChooseBattlers(list []*Item) ([]*Item, []int) {
	fighterOneIndex := rand.Intn(len(list))
	fighterTwoIndex := rand.Intn(len(list))

	for fighterOneIndex == fighterTwoIndex {
		fighterTwoIndex = rand.Intn(len(list))
	}

	fighterOne := list[fighterOneIndex]
	fighterTwo := list[fighterTwoIndex]

	combatants := []*Item{fighterOne, fighterTwo}
	indexes := []int{fighterOneIndex, fighterTwoIndex}
	return combatants, indexes
}
