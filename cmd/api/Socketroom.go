package main

import (
	"fmt"
	"time"
)

type RoomList map[string]Room
type Room struct {
	clientList      ClientList
	owner           *Client
	name            string
	currentQuestion Question
	// gameEnded  bool
}
type Question struct {
	id     int32
	Text   string
	A      string
	B      string
	Answer string
}

func generate(q *Question) {
	q = &Question{
		id:     0,
		Text:   "ss",
		A:      "hello",
		B:      "kol",
		Answer: "B",
	}

}
func (room *Room) startGame() {

	ticker := time.NewTicker(3 * time.Second)
	counter := 0
	fmt.Println("game about to start")
	for range ticker.C {
		//add scores based on past
		generate(&room.currentQuestion)
		counter++
		if counter == 4 {
			fmt.Println("game ended")
			break
		}
		fmt.Println("question: ", counter, " ", room.currentQuestion)
	}
}
func (room *Room) scoreRound() {

}
