package model

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

var (
	config = &bongo.Config{
		ConnectionString: os.Getenv("connectionStr"),
		Database:         "f1",
	}
	connection *bongo.Connection
	connErr    error
)

func init() {
	connection, connErr = bongo.Connect(config)
	if connErr != nil {
		initDb()
		panic(connErr)
	}
}

func initDb() {
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

func GetRace(year, circuitid int64) (Races, error) {
	race := &Races{}
	err := connection.Collection("Races").FindOne(bson.M{"year": year, "circuitid": circuitid}, race)
	return *race, err
}

func GetLaptTimes(raceid, driverid int64) []LapTimes {
	laptime := &LapTimes{}
	laptimes := []LapTimes{}
	results := connection.Collection("LapTimes").Find(bson.M{"raceid": raceid, "driverid": driverid})
	for results.Next(laptime) {
		laptimes = append(laptimes, *laptime)
	}
	return laptimes
}

func GetDriver(driverid int64) (Drivers, error) {
	driver := &Drivers{}
	err := connection.Collection("Drivers").FindOne(bson.M{"driverid": driverid}, driver)
	return *driver, err
}

func GetConstructor(constructorid int64) (Constructors, error) {
	constructor := &Constructors{}
	err := connection.Collection("Constructors").FindOne(bson.M{"constructorid": constructorid}, constructor)
	return *constructor, err
}

func GetResults(raceid, driverid int64) (Results, error) {
	results := &Results{}
	err := connection.Collection("Results").FindOne(bson.M{"raceid": raceid, "driverid": driverid}, results)
	return *results, err
}
