package player

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

// OnlineAgent is a concrete implementation of PlayerAgent for an online client
// Has a WebSocket Connection with the frontend client
type OnlineAgent struct {
	messageChannel chan Message
	playerID       uuid.UUID
	playerLog      *log.Logger
	roomID         uuid.UUID
}

func CreateOnlineAgent() *OnlineAgent {
	newID := uuid.New()
	player := OnlineAgent{
		messageChannel: make(chan Message, 0),
		playerID:       newID,
		playerLog:      log.New(os.Stdout, fmt.Sprintf("Player[#%v]", newID), 1),
	}

	player.playerLog.Println("OnlineAgent created.")

	return &player
}

func (agent *OnlineAgent) RegisterToRoom() {
	agent.playerLog.Printf("Player joined Room")
}
