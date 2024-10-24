package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/demartinom/list-ranker/pkg/battle"
)

func main() {
	list := flag.String("list", "game-data/cities.csv", "list to use for ranking")
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

	itemSlice := battle.ConvertToSlice(listItems)

	var ready string

	fmt.Println("Welcome to List Ranker!")
	fmt.Println("Ready to play? (y/n)")
	fmt.Scanln(&ready)

	switch ready {
	case "y":
		battle.Battle(&itemSlice)
	default:
		break
	}
}
