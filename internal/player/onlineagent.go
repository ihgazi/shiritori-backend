package player

import "fmt"

// OnlineAgent is a concrete implementation of PlayerAgent for an online client
// Has a WebSocket Connection with the frontend client
type OnlineAgent struct {
	messageChannel chan Message
}

func CreateOnlineAgent() *OnlineAgent {
	fmt.Println("OnlineAgent created.")

	player := OnlineAgent{
		messageChannel: make(chan Message, 0),
	}

	return &player
}

func (agent *OnlineAgent) RegisterToRoom() {
	fmt.Println("Player joined Room")
}
