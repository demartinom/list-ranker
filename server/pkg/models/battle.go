package models

import (
	"encoding/csv"
	"fmt"
	"log"
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
