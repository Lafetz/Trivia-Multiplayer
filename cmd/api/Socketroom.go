package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type RoomList map[string]*Room
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

func (room *Room) generateQuestion() {
	room.currentQuestion = Question{
		id:     0,
		Text:   "ss",
		A:      "hello",
		B:      "kol",
		Answer: rand.Int31n(4),
	}

}
func (room *Room) startGame() {

	ticker := time.NewTicker(1 * time.Second)
	counter := 0
	fmt.Println("game about to start")
	for range ticker.C {

		if counter == 3 {
			fmt.Println("game ended")
			room.scoreEvent()
			break
		}
		fmt.Print("count ", counter)
		counter++
		//
		room.roundStart()
		room.sendQuestion()
		///
		//fmt.Println("question: ", counter, " ", room.currentQuestion)
	}
}

func (room *Room) roundStart() {
	room.generateQuestion()
	room.restUserAnswer()

}
func (room *Room) restUserAnswer() {
	for c := range room.clientList {
		c.answer = 0
	}
}
func (room *Room) sendQuestion() {
	payload, err := json.Marshal(room.currentQuestion)
	if err != nil {
		fmt.Println(err)
	}
	event := Event{
		Type:    EventSendQuestion,
		Payload: payload,
	}
	for c := range room.clientList {
		c.egress <- event
	}
}

// func (room *Room) sendFinalScores() {

// 	for c := range room.clientList {

//			fmt.Print(c.score, " ")
//		}
//	}
func (room *Room) scoreEvent() {
	array := []UserScore{}
	for c := range room.clientList {
		array = append(array, UserScore{
			name:  c.name,
			score: int(c.score),
		})
	}
	type myJSON struct {
		Array []UserScore
	}
	fmt.Println(array)
	payload, err := json.Marshal(myJSON{Array: array})
	if err != nil {
		fmt.Println(err)
	}
	event := Event{
		Type:    EventFinalScores,
		Payload: payload,
	}
	for c := range room.clientList {
		c.egress <- event
	}
}

// func (room *Room) sendFinalScores() {
// 	for c := range room.clientList {
// 		fmt.Print(c.score, " ")
// 	}
// }
