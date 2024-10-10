package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"

	"github.com/demartinom/list-ranker/battle"
	"github.com/demartinom/list-ranker/filehandler"
)

func main() {
	list := flag.String("list", "cities.csv", "list to use for ranking")
	flag.Parse()

	file, err := os.Open(*list)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	// Read header line so it is not passed into convert function
	reader.Read()

	listItems, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	itemSlice := filehandler.ConvertToSlice(listItems)

	battle.Battle(itemSlice)
}
