package main

import (
	"encoding/json"
	"github.com/magleff/galaxy-challenge/core"
	"github.com/magleff/galaxy-challenge/models"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		universe models.Universe
		response models.Response
	)

	if err = json.NewDecoder(r.Body).Decode(&universe); err != nil {
		log.Println("An error occured while decoding json:", err)
	} else {
		log.Println("Message received for turn", universe.Config.Turn)
	}

	response, err = core.Analyse(universe)

	if err != nil {
		log.Println("An error occured while analysing the current state of the game:", err)
	} else {
		log.Println("Sending back data for turn", universe.Config.Turn)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	log.Println("Running the server on port 3000")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
