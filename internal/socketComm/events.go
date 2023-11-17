package socketComm

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type UserEvent struct {
	event Event
	from  string
}
type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
type EventHandler func(event Event, c *Client) error

const (
	//sent from user to server
	EventSendMessage = "send_message"
	EventSendAnswer  = "send_answer"
	//from server
	EventError        = "error"
	EventGameStart    = "game_start"
	EventSendQuestion = "send_Question"
	EventFinalScores  = "final_scores"
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

func SendMessage(event Event, c *Client) error {

	forwardPayload(event, c.room)
	fmt.Println(event.Type, string(event.Payload))
	return nil
}

func SendAnswr(event Event, c *Client) error {
	fmt.Println("Event send answer")
	if c.answer != 0 {
		fmt.Println("user submited again")
		return nil

	}
	userAnswer, err := strconv.Atoi(string(event.Payload))
	fmt.Println("payload", event.Payload, "is ", userAnswer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//fmt.Println("test", userAnswer, " ", c.room.currentQuestion.Answer)
	if int32(userAnswer) == c.room.currentQuestion.Answer {
		c.score++
	}
	c.answer = int32(userAnswer)
	forwardPayload(event, c.room)
	return nil
}
func StartGame(event Event, c *Client) error {
	fmt.Println("start game")
	go c.room.startGame()
	return nil
}

func SendScores(room *Room) {
	array := []UserScore{}
	for c := range room.clientList {
		array = append(array, UserScore{
			Name:  c.name,
			Score: int(c.score),
		})
	}
	event := createEvent(EventFinalScores, &ScoresList{Scores: array})
	forwardPayload(event, room)
}

func sendQuestion(room *Room) {
	event := createEvent(EventSendQuestion, room.currentQuestion)
	forwardPayload(event, room)
}

// helper fucs
func createEvent(eventType string, payload interface{}) Event {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	return Event{
		Type:    eventType,
		Payload: payloadJson,
	}
}
func forwardPayload(event Event, room *Room) error { //sends payload to everyone in a room
	for c := range room.clientList {
		c.egress <- event
	}
	return nil
}
