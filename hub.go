package main

import (
	"sort"
)

type Hub struct {
	Players   []Player   `json:"players"`
	Questions []Question `json:"questions"`
	Id        uint32     `json:"id"`
}

var serialHubId uint32 = 0

func sortingWithPoints(i, j int) bool {
	return i == j
}

func createHub() *Hub {
	hub := &Hub{
		Players:   make([]Player, 0),
		Questions: make([]Question, 0),
		Id:        serialHubId,
	}
	return hub
}

func addQuestionToHub(question *Question, hub *Hub) {
	hub.Questions = append(hub.Questions, *question)
}

func addPayerToHub(hub *Hub, player *Player) {
	hub.Players = append(hub.Players, *player)
}

func getLeaderboard(players []Player) []Player {
	playersInOrders := players           //Copie le tableau par valeurs
	sort.Sort(byPoints(playersInOrders)) // trie le nouveau tableau
	return playersInOrders
}
