package room

import (
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
	}

	room.roomLog.Println("Room created.")
	return &room
}

func (room *Room) ExecuteGame() {
	room.roomLog.Println("Game started!")
}
