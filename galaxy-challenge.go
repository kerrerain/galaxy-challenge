package main

import (
	"encoding/json"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"github.com/magleff/galaxy-challenge/ias/leraje"
	"log"
	"net/http"
)

var G *game.Map

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		status dto.Status
		move   dto.Move
	)

	if err = json.NewDecoder(r.Body).Decode(&status); err != nil {
		log.Println("An error occured while decoding json:", err)
	} else {
		log.Println("Message received for turn", status.Config.Turn)
	}

	move, err = makeMove(status)

	if err != nil {
		log.Println("An error occured while analysing the current state of the game:", err)
	} else {
		log.Println("Sending back data for turn", status.Config.Turn)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(move)
}

func main() {
	log.Println("Running the server on port 3000")

	G = &game.Map{}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func makeMove(status dto.Status) (dto.Move, error) {
	updateGame(status)
	return paimon.Run(G), nil
}

func updateGame(status dto.Status) {
	if status.Config.Turn > G.Turn {
		G.Update(status)
	} else {
		G = &game.Map{}
		G.Update(status)
		G.InitDistanceMap()
	}
}
