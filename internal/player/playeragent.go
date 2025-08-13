package player

import (
	"encoding/json"

	"github.com/google/uuid"
)

const (
	MessageTypeMove   = "move"
	MessageTypeError  = "error"
	MessageTypeSystem = "system"
)

// Generic message that can be used to indicate the move played
// or an update to the game state
type Message struct {
	Type      string          `json:"type"`
	ID        string          `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Sender    string          `json:"sender"`
	Payload   json.RawMessage `json:"payload"`
}

type MovePayload struct {
	Word string `json:"word"` // The word played by the player
}

type ErrorPayload struct {
	Code    int    `json:"code"`    // Error code
	Message string `json:"message"` // Error message
}

type SystemPayload struct {
	Message string `json:"message"` // System message
}

// PlayerAgent represents an entity in an active game
// The agent can be of three states
// - queued for matchmaking
// - waiting for opponent's move
// - waiting to play own move
//
// RegisterToRoom() - once MatchMaker allocates a room
// send room details to client to update frontend
type PlayerAgent interface {
	RegisterToRoom(roomID uuid.UUID)
	WriteMessage(message *Message) error
	ReadMessage() (*Message, error)
	CloseAgent()
}
