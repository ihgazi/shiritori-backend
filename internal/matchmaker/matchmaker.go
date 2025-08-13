package matchmaker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ihgazi/shiritori/internal/player"
	"github.com/ihgazi/shiritori/internal/room"
)
import "sync"

// A matchmaker that receieves queue requests and attemps to allocate a Room
// TODO: implement a seperate queueing logic, for complex matchmaking
type MatchMaker struct {
	playerQueue []player.PlayerAgent
	queueLock   sync.Mutex
	TestMode    bool
}

// Add a player to queue
// If another player exists in queue remove both players from queue and start a game
// returns true if a room was allocated else false
func (matcher *MatchMaker) QueuePlayer(currPlayer player.PlayerAgent) bool {
	matcher.queueLock.Lock()
	defer matcher.queueLock.Unlock()

	queueLength := len(matcher.playerQueue)

	if queueLength == 0 {
		// wait for another player to queue in
		matcher.playerQueue = append(matcher.playerQueue, currPlayer)
		return false
	}

	oppPlayer := matcher.playerQueue[queueLength-1]
	matcher.playerQueue = matcher.playerQueue[:queueLength-1]

	currRoom := room.MakeRoom(currPlayer, oppPlayer)
	currPlayer.RegisterToRoom(currRoom.GetID())
	oppPlayer.RegisterToRoom(currRoom.GetID())

	if matcher.TestMode != true {
		go currRoom.ExecuteGame()
	}
	return true
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// HTTP Handler for queueing
// Upgrades to Websocket Connection
// Instantiates PlayerAgent and calls QueuePlayer()
func (matcher *MatchMaker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Player queued")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	currPlayer := player.CreateOnlineAgent(conn)
	matcher.QueuePlayer(currPlayer)
}
