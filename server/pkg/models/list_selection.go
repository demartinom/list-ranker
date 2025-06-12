package models

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func PremadeList() []string {
	files := getFileNames()
	return files
}

func getFileNames() []string {
	dirPath := "game-data"
	var files []string
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fileName := d.Name()
			files = append(files, fileName[:len(fileName)-4])
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

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
		itemsList = append(itemsList, &Item{Name: itemInput[0], Score: 0, Rounds: 0})
	}

	return itemsList
}
