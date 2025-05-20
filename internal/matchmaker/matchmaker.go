package matchmaker

import (
	"net/http"

	"github.com/ihgazi/shiritori/internal/player"
	"github.com/ihgazi/shiritori/internal/room"
)
import "sync"

// A matchmaker that receieves queue requests and attemps to allocate a Room
// TODO: implement a seperate queueing logic, for complex matchmaking
type MatchMaker struct {
	playerQueue []player.PlayerAgent
	queueLock   sync.Mutex
}

func (matcher *MatchMaker) QueuePlayer() bool {
	matcher.queueLock.Lock()
	defer matcher.queueLock.Unlock()

	currPlayer := player.CreateOnlineAgent()

	queueLength := len(matcher.playerQueue)

	if queueLength == 0 {
		// wait for another player to queue in
		matcher.playerQueue = append(matcher.playerQueue, currPlayer)
		return false
	}

	oppPlayer := matcher.playerQueue[queueLength-1]
	matcher.playerQueue = matcher.playerQueue[:queueLength-1]

	currRoom := room.MakeRoom(currPlayer, oppPlayer)
	go currRoom.ExecuteGame()
	return true
}

func (matcher *MatchMaker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if matcher.QueuePlayer() == true {
		w.Write([]byte("Match found!"))
	} else {
		w.Write([]byte("Waiting for other players..."))
	}
}
