package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuaj1397/f1-data-visualization/api"
)

var port = "3005"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/Races/{year}/{circuitid}", api.GetRace).Methods("GET")
	router.HandleFunc("/LapTimes/{raceid}/{driverid}", api.GetRaceLapTimes).Methods("GET")
	router.HandleFunc("/Drivers/{driverid}", api.GetDriver).Methods("GET")
	router.HandleFunc("/Constructors/{constructorid}", api.GetConstructor).Methods("GET")
	router.HandleFunc("/Results/{raceid}/{driverid}", api.GetResults).Methods("GET")

	fmt.Println("Now listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
