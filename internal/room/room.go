package room

import (
	"fmt"
	"github.com/ihgazi/shiritori/internal/player"
)

// Represents an active game in progress
type Room struct {
	// Store reference to the two PlayerAgents in game
	roomPlayers []player.PlayerAgent
}

func MakeRoom(player1 player.PlayerAgent, player2 player.PlayerAgent) *Room {
	playerSlice := make([]player.PlayerAgent, 2)
	playerSlice[0] = player1
	playerSlice[1] = player2

	room := Room{
		roomPlayers: playerSlice,
	}

	fmt.Println("Room created.")
	return &room
}

func (room *Room) ExecuteGame() {
	fmt.Println("Game started!")
}
