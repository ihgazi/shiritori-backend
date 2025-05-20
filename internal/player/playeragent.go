package player

import (
	"github.com/google/uuid"
)

// Generic message that can be used to indicate the move played
// or an update to the game state
type Message struct {
	MessageType  string `json:"type"`
	MessageValue string `json:"value"`
	MessageID    string `json:"id"`
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
}
