package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuaj1397/f1-data-visualization/api"
	"github.com/joshuaj1397/f1-data-visualization/model"
)

var port = "3005"
var recreateDb = true

func main() {
	if recreateDb {
		model.InitDb()
	}
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./frontend/"))

	// API
	router.HandleFunc("/Circuit/{circuitId}", api.GetCircuit).Methods("GET")
	router.HandleFunc("/Races/{year}/{circuitId}", api.GetRace).Methods("GET")
	router.HandleFunc("/LapTimes/{raceId}/{driverId}", api.GetRaceLapTimes).Methods("GET")
	router.HandleFunc("/Drivers/{driverId}", api.GetDriver).Methods("GET")
	router.HandleFunc("/Constructors/{constructorId}", api.GetConstructor).Methods("GET")
	router.HandleFunc("/Results/{raceId}/{driverId}", api.GetResults).Methods("GET")

	// Webpages
	router.PathPrefix("/").Handler(fs)

	fmt.Println("Now listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
