package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
	"github.com/magleff/galaxy-challenge/ias/agares"
	"github.com/magleff/galaxy-challenge/ias/amon"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var Games map[int16]*game.Map
var TurnsLog map[int16]*dto.TurnLog

const FIRST_PLAYER_AI = agares.NAME
const SECOND_PLAYER_AI = amon.NAME
const SOLO_MODE_MAP = "criticorum"

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

	start := time.Now()
	move = makeMove(status, 1, FIRST_PLAYER_AI)
	elapsed := time.Since(start)

	log.Printf("Took %s", elapsed)
	log.Println("Sending back data for turn", status.Config.Turn, move)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(move)

	logToFile(status, move)
}

func logToFile(status dto.Status, move dto.Move) {
	gameID := status.Config.ID

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	f, err := os.Create("logs/agares_" + strconv.Itoa(int(gameID)) + ".json")
	defer f.Close()

	if err != nil {
		log.Panic(err)
	}

	if TurnsLog[gameID] == nil {
		TurnsLog[gameID] = &dto.TurnLog{
			Data: make([]interface{}, 0),
		}
	}

	TurnsLog[gameID].Data = append(TurnsLog[gameID].Data, []interface{}{
		status,
		move,
	})

	w := bufio.NewWriter(f)
	json.NewEncoder(w).Encode(TurnsLog[gameID].Data)

	w.Flush()
}

func main() {
	soloModeFlag := flag.Bool("s", false, "Solo mode game")
	flag.Parse()

	Games = make(map[int16]*game.Map)
	TurnsLog = make(map[int16]*dto.TurnLog)

	if *soloModeFlag {
		soloModeHandler()
	} else {
		log.Println("Running the server on port 80")

		http.HandleFunc("/", classicModeHandler)
		http.ListenAndServe(":80", nil)
	}
}

func soloModeHandler() {
	log.Println("Playing in solo mode")

	var (
		err     error
		planets []dto.StatusPlanet
	)

	file, e := ioutil.ReadFile("./maps/" + SOLO_MODE_MAP + ".json")
	if e != nil {
		log.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	if err = json.Unmarshal(file, &planets); err != nil {
		log.Println("An error occured while decoding json:", err)
	} else {
		log.Printf("Loaded map: %s", SOLO_MODE_MAP)
	}

	gameMap := &game.Map{}
	gameMap.Update(dto.Status{
		Planets: planets,
	})
	gameMap.InitDistanceMap()

	timeline := engine.CreateTimeline(gameMap)

	for i := 1; i <= 200; i++ {
		log.Printf("Turn %d", i)
		status := timeline.Status()

		firstPlayerMove := makeMove(status, 1, FIRST_PLAYER_AI)
		timeline.ScheduleMoveForNextTurn(1, firstPlayerMove)
		logToFile(status, firstPlayerMove)

		secondPlayerMove := makeMove(status, 2, SECOND_PLAYER_AI)
		timeline.ScheduleMoveForNextTurn(2, secondPlayerMove)

		timeline.NextTurn()
	}
}

func classicModeHandler(w http.ResponseWriter, r *http.Request) {
	handler(w, r)
}

func makeMove(status dto.Status, playerID int16, aiName string) dto.Move {
	updateGame(status)

	var move dto.Move

	switch aiName {
	case agares.NAME:
		move = agares.Run(Games[status.Config.ID], playerID)
	default:
		move = amon.Run(Games[status.Config.ID], playerID)
	}

	return move
}

func updateGame(status dto.Status) {
	if Games[status.Config.ID] != nil {
		Games[status.Config.ID].Update(status)
	} else {
		Games[status.Config.ID] = &game.Map{
			ID: status.Config.ID,
		}
		Games[status.Config.ID].Update(status)
		Games[status.Config.ID].InitDistanceMap()
	}
}
