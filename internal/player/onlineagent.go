package player

import (
	"encoding/json"
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

// Add player to a room and inform client
func (agent *OnlineAgent) RegisterToRoom(roomID uuid.UUID) {
	agent.playerLog.Printf("Player joined Room")

	agent.roomID = roomID
	payloadBytes, _ := json.Marshal(SystemPayload{
		Message: "REGISTER",
	})

	agent.WriteMessage(&Message{
		Type:    MessageTypeSystem,
		Sender:  roomID.String(),
		Payload: payloadBytes,
	})
}

// Write a message to the client
func (agent *OnlineAgent) WriteMessage(message *Message) error {
	if err := agent.playerConn.WriteJSON(message); err != nil {
		return err
	}

	return nil
}

// Read message from websocket
func (agent *OnlineAgent) ReadMessage() (*Message, error) {
	message := new(Message)
	err := agent.playerConn.ReadJSON(message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (agent *OnlineAgent) CloseAgent() {
	agent.playerConn.Close()
}
