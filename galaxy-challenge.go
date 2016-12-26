package main

import (
	"encoding/json"
	"github.com/magleff/galaxy-challenge/base"
	"github.com/magleff/galaxy-challenge/models"
	"log"
	"net/http"
)

var game models.Game

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		request  models.Request
		response models.Response
	)

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("An error occured while decoding json:", err)
	} else {
		log.Println("Message received for turn", request.Config.Turn)
	}

	identifyGame(request)

	response, err = base.Analyse(game)

	if err != nil {
		log.Println("An error occured while analysing the current state of the game:", err)
	} else {
		log.Println("Sending back data for turn", request.Config.Turn)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	log.Println("Running the server on port 3000")

	game = models.CreateNewGame()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func identifyGame(request models.Request) {
	if request.Config.Turn > game.Turn {
		game.Update(request)
	} else {
		game = models.CreateNewGame()
	}
}
