package room

import (
	"slices"
)

// Dictionary implements the game logic, validating and recording each word sent by a player
// TODO: Implement faster logic using Trie data structure
type dictionary struct {
	moveHistory []string
}

func MakeDictionary() *dictionary {
	return &dictionary{
		moveHistory: make([]string, 0),
	}
}

func (dict *dictionary) Record(currWord string) bool {
	if slices.Contains(dict.moveHistory, currWord) {
		return false
	}

	dict.moveHistory = append(dict.moveHistory, currWord)
	return true
}
