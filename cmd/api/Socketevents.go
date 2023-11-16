package main

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
type EventHandler func(event Event, c *Client) error

const (
	EventSendMessage = "send_message"
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
	fmt.Println(event.Type, string(event.Payload))
	return nil
}
