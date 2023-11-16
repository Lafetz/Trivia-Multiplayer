package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
type EventHandler func(event Event, c *Client) error

const (
	EventSendMessage = "send_message"
	EventSendAnswer  = "send_answer"
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

func SendMessage(event Event, c *Client) error {
	for k, v := range c.manager.roomList {
		if k == c.room.name {
			for k, _ := range v.clientList {
				k.egress <- event
			}
		}
	}
	// if event.Type == EventUserAnswer {
	// 	if c.score != 0 {
	// 		fmt.Println("user submited again")
	// 		return nil

	// 	}
	// 	userAnswer, err := strconv.Atoi(string(event.Payload))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	if int32(userAnswer) == c.room.currentQuestion.Answer {
	// 		c.score++
	// 	}

	// }
	fmt.Println(event.Type, string(event.Payload))
	return nil
}

func SendAnswr(event Event, c *Client) error {
	fmt.Println("Event send answer")
	if c.score != 0 {
		fmt.Println("user submited again")
		return nil

	}
	userAnswer, err := strconv.Atoi(string(event.Payload))
	if err != nil {
		fmt.Println(err)
	}
	if int32(userAnswer) == c.room.currentQuestion.Answer {
		c.score++
	}
	return nil
}
