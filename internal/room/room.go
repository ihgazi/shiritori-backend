package room

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/ihgazi/shiritori/internal/player"
)

// Represents an active game in progress
type Room struct {
	// Store reference to the two PlayerAgents in game
	roomPlayers []player.PlayerAgent
	roomID      uuid.UUID
	roomLog     *log.Logger
	roomDict    *dictionary
}

func MakeRoom(player1 player.PlayerAgent, player2 player.PlayerAgent) *Room {
	playerSlice := make([]player.PlayerAgent, 2)
	playerSlice[0] = player1
	playerSlice[1] = player2

	newID := uuid.New()

	room := Room{
		roomPlayers: playerSlice,
		roomID:      newID,
		roomLog:     log.New(os.Stdout, fmt.Sprintf("Room[%v]", newID), 1),
		roomDict:    MakeDictionary(),
	}

	room.roomLog.Println("Room created.")
	return &room
}

func (room *Room) makeMessage(value string) *player.Message {
	payloadBytes, _ := json.Marshal(player.SystemPayload{
		Message: value,
	})

	return &player.Message{
		Type:    player.MessageTypeSystem,
		Sender:  room.roomID.String(),
		Payload: payloadBytes,
	}
}

func (room *Room) makeMoveMessage(value string) *player.Message {
	payloadBytes, _ := json.Marshal(player.MovePayload{
		Word: value,
	})

	return &player.Message{
		Type:    player.MessageTypeMove,
		Sender:  room.roomID.String(),
		Payload: payloadBytes,
	}
}

func (room *Room) ExecuteGame() {
	room.roomLog.Println("Game started!")
	defer room.roomPlayers[0].CloseAgent()
	defer room.roomPlayers[1].CloseAgent()

	room.roomPlayers[0].WriteMessage(room.makeMessage("PLAY"))
	room.roomPlayers[1].WriteMessage(room.makeMessage("WAIT"))

	turn := 0

	for {
		userMove, err := room.roomPlayers[turn].ReadMessage()
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		var move player.MovePayload
		err = json.Unmarshal(userMove.Payload, &move)
		if err != nil {
			room.roomLog.Printf("Error unmarshalling move: %s", err.Error())
			return
		}

		if room.roomDict.Record(move.Word) == false {
			room.roomPlayers[turn].WriteMessage(room.makeMessage("INVALID"))
		} else {
			turn = (turn + 1) % 2
			room.roomPlayers[turn].WriteMessage(room.makeMoveMessage(move.Word))
		}
	}
}

func (room *Room) GetID() uuid.UUID {
	return room.roomID
}
