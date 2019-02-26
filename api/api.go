package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/joshuaj1397/f1-data-visualization/model"
)

func GetRace(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRace called")
	params := mux.Vars(r)
	year, err := strconv.ParseInt(params["year"], 10, 32)
	circuitid, err := strconv.ParseInt(params["circuitid"], 10, 32)
	race, err := model.GetRace(year, circuitid)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(race)
}

func GetRaceLapTimes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raceid, err := strconv.ParseInt(params["raceid"], 10, 32)
	driverid, err := strconv.ParseInt(params["driverid"], 10, 32)
	laptimes := model.GetLaptTimes(raceid, driverid)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(laptimes)
}

func GetDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	driverid, err := strconv.ParseInt(params["year"], 10, 32)
	driver, err := model.GetDriver(driverid)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(driver)
}

func GetConstructor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	constructorid, err := strconv.ParseInt(params["constructorid"], 10, 32)
	constructor, err := model.GetConstructor(constructorid)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(constructor)
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raceid, err := strconv.ParseInt(params["raceid"], 10, 32)
	driverid, err := strconv.ParseInt(params["driverid"], 10, 32)
	results, err := model.GetResults(raceid, driverid)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(results)
}
