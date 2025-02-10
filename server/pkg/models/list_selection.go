package models

import (
	"io/fs"
	"log"
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