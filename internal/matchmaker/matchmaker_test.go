package matchmaker

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ihgazi/shiritori/internal/player"
	"github.com/stretchr/testify/assert"
)

type mockPlayer struct {
	registeredRoomID uuid.UUID
}

func (m *mockPlayer) RegisterToRoom(roomID uuid.UUID) {
	m.registeredRoomID = roomID
}

func (m *mockPlayer) WriteMessage(msg *player.Message) error {
	return nil
}

func (m *mockPlayer) ReadMessage() (*player.Message, error) {
	return nil, nil
}

func (m *mockPlayer) CloseAgent() {}

func TestQueuePlayer_WaitsWhenAlone(t *testing.T) {
	matcher := &MatchMaker{}
	p1 := &mockPlayer{}

	roomAllocated := matcher.QueuePlayer(p1)

	assert.False(t, roomAllocated)
	assert.Equal(t, 1, len(matcher.playerQueue))
	assert.Equal(t, p1, matcher.playerQueue[0])
}

func TestQueuePlayer_StartsGameWhenPairFound(t *testing.T) {
	matcher := &MatchMaker{TestMode: true}

	p1 := &mockPlayer{}
	p2 := &mockPlayer{}

	matcher.QueuePlayer(p1)
	roomAllocated := matcher.QueuePlayer(p2)

	assert.True(t, roomAllocated)
	assert.Equal(t, 0, len(matcher.playerQueue))
	assert.NotEqual(t, uuid.Nil, p1.registeredRoomID)
	assert.Equal(t, p1.registeredRoomID, p2.registeredRoomID)
}
