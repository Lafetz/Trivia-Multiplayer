package main

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
	EventGameStart    = "game_start"
	EventSendQuestion = "send_Question"
	EventFinalScores  = "final_scores"
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

func SendMessage(event Event, c *Client) error {
	// for k, v := range c.manager.roomList {
	// 	if k == c.room.name {
	// 		for k, _ := range v.clientList {
	// 			k.egress <- event
	// 		}
	// 	}
	// }
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
	fmt.Println("test", userAnswer, " ", c.room.currentQuestion.Answer)
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
func forwardPayload(event Event, room *Room) error { //sends payload to everyone in a room
	for k := range room.clientList {
		k.egress <- event
	}
	return nil
}

type UserScore struct {
	name  string
	score int
}

// func createEvent(eventType string, payload interface{}) *Event {
// 	payloadJson, err := json.Marshal(payload)
// 	return &Event{
// 		Type:    eventType,
// 		Payload: payload,
// 	}
// }
