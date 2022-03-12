package main

import (
	"errors"
)

var ErrorPlayerIdNotFound error = errors.New("Player Id Not Found")

var serialPlayerId uint32 = 1

var players []Player

type Player struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       uint32 `json:"id"`
	Points   int    `json:"points"`
}

type byPoints []Player

func (player byPoints) Len() int           { return len(player) }
func (player byPoints) Less(i, j int) bool { return player[i].Points < player[j].Points }
func (player byPoints) Swap(i, j int)      { player[i], player[j] = player[j], player[i] }

func createPlayer(username string, password string) *Player {
	player := &Player{
		Username: username,
		Password: password,
		Id:       serialPlayerId,
		Points:   0,
	}
	serialPlayerId += 1
	return player
}

func addPlayer(player *Player) {
	players = append(players, *player)
}

func removePlayer(id uint32) error {
	for i, player := range players {
		if player.Id == id {
			players = append(players[:i], players[i+1:]...)
			return nil
		}
	}
	return ErrorPlayerIdNotFound
}

func getPlayer(id uint32) (*Player, error) {
	for _, player := range players {
		if player.Id == id {
			return &player, nil
		}
	}
	return nil, ErrorPlayerIdNotFound
}

func updatePlayer(id uint32, username string) error {
	for _, player := range players {
		if player.Id == id {
			player.Username = username
			return nil
		}
	}
	return ErrorPlayerIdNotFound
}

func getPlayers() []Player {
	return players
}

func verifyPlayer(username string, password string) (*Player, error) {
	for _, player := range players {
		if player.Username == username && player.Password == password {
			return &player, nil
		}
	}
	return nil, errors.New("No user with theses credentials")
}
