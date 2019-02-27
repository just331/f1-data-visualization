package model

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectionStr = os.Getenv("connectionStr")
	dbName        = "f1"
	ctx, _        = context.WithTimeout(context.Background(), 10*time.Second)
	client        *mongo.Client
	db            *mongo.Database
)

func init() {
	client, _ = mongo.NewClient(options.Client().ApplyURI(connectionStr))
	connErr := client.Connect(ctx)
	db = client.Database(dbName)

	if connErr != nil {
		panic(connErr)
	}

	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
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

	for _, file := range files {
		importCSVtoMongoDb(file)
	}

}

func importCSVtoMongoDb(file string) error {
	reg := regexp.MustCompile(`[\\/].*`)
	match := reg.FindStringSubmatch(file)
	fileName := match[0]
	fileName = fileName[1 : len(fileName)-4]
	collection := strings.Title(fileName)
	cmd := exec.Command("mongoimport", "--db", "f1", "--collection", collection, "--type", "csv", "--headerline", "--file", file)
	err := cmd.Run()
	return err
}

func GetCircuit(circuitId int) (Circuits, error) {
	circuit := &Circuits{}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err := db.Collection("Circuits").FindOne(ctx, bson.M{"circuitId": circuitId}).Decode(circuit)
	return *circuit, err
}

func GetRace(year, circuitId int) (Races, error) {
	race := &Races{}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err := db.Collection("Races").FindOne(ctx, bson.M{"year": year, "circuitId": circuitId}).Decode(race)
	return *race, err
}

func GetLapTimes(raceId, driverId int) []LapTimes {
	laptime := &LapTimes{}
	laptimes := []LapTimes{}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := db.Collection("LapTimes").Find(ctx, bson.M{"raceId": raceId, "driverId": driverId})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var elem LapTimes
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		laptimes = append(laptimes, *laptime)
	}
	return laptimes
}

func GetDriver(driverId int) (Drivers, error) {
	driver := &Drivers{}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err := db.Collection("Drivers").FindOne(ctx, bson.M{"driverId": driverId}).Decode(driver)
	return *driver, err
}

func GetConstructor(constructorId int) (Constructors, error) {
	constructor := &Constructors{}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err := db.Collection("Constructors").FindOne(ctx, bson.M{"constructorId": constructorId}).Decode(constructor)
	return *constructor, err
}

func GetResults(raceId, driverId int) (Results, error) {
	results := &Results{}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err := db.Collection("Results").FindOne(ctx, bson.M{"raceId": raceId, "driverId": driverId}).Decode(results)
	return *results, err
}
