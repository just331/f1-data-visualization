package model

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-bongo/bongo"
)

var (
	config = &bongo.Config{
		ConnectionString: os.Getenv("connectionStr"),
		Database:         "f1",
	}
	connection bongo.Connection
)

// func init() {
// 	connection, err := bongo.Connect(config)
// 	if err != nil {
// 		panic(err)
// 	}
// }

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

	for _, file := range files {
		importCSVtoMongoDb(file)
	}

}

func importCSVtoMongoDb(file string) error {
	reg := regexp.MustCompile("\\/.*")
	match := reg.FindStringSubmatch(file)
	fileName := match[0]
	fileName = fileName[1 : len(fileName)-4]
	collection := strings.Title(fileName)
	cmd := exec.Command("mongoimport", "--db", "f1", "--collection", collection, "--type", "csv", "--headerline", "--file", file)
	err := cmd.Run()
	return err
}
