package main

/*import (
	"encoding/json"
	"github.com/magleff/galaxy-challenge/core"
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/move"
	"github.com/magleff/galaxy-challenge/models/status"
	"log"
	"net/http"
)

var G *game.Game

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		status status.Game
		move   *move.Move
	)

	if err = json.NewDecoder(r.Body).Decode(&status); err != nil {
		log.Println("An error occured while decoding json:", err)
	} else {
		log.Println("Message received for turn", status.Config.Turn)
	}

	move, err = makeMove(status)

	log.Println(len(move.Fleets), "fleets sent.")

	if err != nil {
		log.Println("An error occured while analysing the current state of the game:", err)
	} else {
		log.Println("Sending back data for turn", status.Config.Turn)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*move)
}

func main() {
	log.Println("Running the server on port 3000")

	G = game.CreateNewGame()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func makeMove(status status.Game) (*move.Move, error) {
	updateGame(status)
	return core.MakeMove(G)
}

func updateGame(status status.Game) {
	if status.Config.Turn > G.Turn {
		core.UpdateGame(status, G)
	} else {
		G = game.CreateNewGame()
	}
}*/
