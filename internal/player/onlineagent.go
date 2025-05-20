package player

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// OnlineAgent is a concrete implementation of PlayerAgent for an online client
// Has a WebSocket Connection with the frontend client
type OnlineAgent struct {
	messageChannel chan Message
	playerID       uuid.UUID
	playerLog      *log.Logger
	playerConn     *websocket.Conn
	roomID         uuid.UUID
}

func CreateOnlineAgent(conn *websocket.Conn) *OnlineAgent {
	newID := uuid.New()
	player := OnlineAgent{
		messageChannel: make(chan Message, 0),
		playerID:       newID,
		playerLog:      log.New(os.Stdout, fmt.Sprintf("Player[#%v]", newID), 1),
		playerConn:     conn,
	}

	player.playerLog.Println("OnlineAgent created.")

	return &player
}

func (agent *OnlineAgent) RegisterToRoom(roomID uuid.UUID) {
	agent.playerLog.Printf("Player joined Room")

	agent.roomID = roomID
	agent.playerConn.WriteJSON(Message{
		MessageType:  "system",
		MessageID:    roomID.String(),
		MessageValue: "found",
	})
}
