package model

import (
	"os"
	"path/filepath"

	"github.com/go-bongo/bongo"
)

var (
	config = &bongo.Config{
		ConnectionString: os.Getenv("connectionStr"),
		Database:         "f1",
	}
	connection bongo.Connection
)

func init() {
	connection, err := bongo.Connect(config)
	if err != nil {
		panic(err)
	}
}

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

func importCSVtoMongoDb() {

}
