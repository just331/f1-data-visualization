package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuaj1397/f1-data-visualization/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/Races/{year}/{circuitid}", api.GetRace).Methods("GET")
	router.HandleFunc("/LapTimes/{raceid}/{driverid}", api.GetRaceLapTimes).Methods("GET")
	router.HandleFunc("/Drivers/{driverid}", api.GetDriver).Methods("GET")
	router.HandleFunc("/Constructors/{constructorid}", api.GetConstructor).Methods("GET")
	router.HandleFunc("/Results/{raceid}/{driverid}", api.GetResults).Methods("GET")

	log.Fatal(http.ListenAndServe(":3005", router))
}
