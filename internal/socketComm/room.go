package socketComm

import (
	"fmt"
	"math/rand"
	"time"
)

type RoomList map[string]*Room
type Room struct {
	clientList      ClientList
	owner           string
	name            string
	CurrentQuestion Question
	// gameEnded  bool
}
type Question struct {
	Id     int32
	Text   string
	A      string
	B      string
	Answer int32
}

func (room *Room) generateQuestion() {
	room.CurrentQuestion = Question{
		Id:     0,
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
			SendScores(room)
			break
		}
		fmt.Print("count ", counter)
		counter++
		//
		room.roundStart()
		sendQuestion(room)
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
func NewRoom(name string) *Room {
	return &Room{
		clientList: make(ClientList),
		owner:      "abel",
		name:       name,
	}
}

type UserScore struct {
	Name  string
	Score int
}
type ScoresList struct {
	Scores []UserScore
}
