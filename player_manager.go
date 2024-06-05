package main

type PlayerManager struct {
	players       []*Player
	currPlayerIdx int
}

func NewPlayerManager(players []*Player) *PlayerManager {
	return &PlayerManager{
		players:       players,
		currPlayerIdx: 0,
	}
}

func (m *PlayerManager) CurrentPlayer() *Player {
	return m.players[m.currPlayerIdx]
}

func (m *PlayerManager) rotate() {
	if m.currPlayerIdx == len(m.players)-1 {
		m.currPlayerIdx = 0
	} else {
		m.currPlayerIdx++
	}
}
