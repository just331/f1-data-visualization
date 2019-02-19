package model

import (
	"os"
	"path/filepath"
)

func InitDb() {
	var files []string
	root := "dataset"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

}
