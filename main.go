package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
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

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range records {
		fmt.Println(v[0])
	}
}
