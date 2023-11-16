package main

import (
	"fmt"
	"time"
)

type RoomList map[string]Room
type Room struct {
	clientList      ClientList
	owner           string
	name            string
	currentQuestion Question
	// gameEnded  bool
}
type Question struct {
	id     int32
	Text   string
	A      string
	B      string
	Answer int32
}

func generate(q *Question) {
	q = &Question{
		id:     0,
		Text:   "ss",
		A:      "hello",
		B:      "kol",
		Answer: 2,
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

// func (room *Room) scoreRound() {
// 	for c := range room.clientList {
// 		if c.answer == 0 {

//			}
//		}
// //	}
// func (room *Room) addUserAnswer(c *Client,answer ) {

// }
