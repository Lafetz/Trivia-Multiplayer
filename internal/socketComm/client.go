package socketComm

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[*Client]bool
type Client struct {
	connection *websocket.Conn
	manager    *Manager
	room       *Room
	score      int32
	answer     int32
	egress     chan Event
	name       string
}

func NewClient(conn *websocket.Conn, manager *Manager, room *Room) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		room:       room,
		egress:     make(chan Event),
	}
}
func (c *Client) SendMessage() { //from server to client
	defer func() {
		c.manager.removeClient(c, c.room.name)
	}()
	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					fmt.Println(err)
				}
				return

			}
			data, err := json.Marshal(message)
			if err != nil {
				fmt.Println(err)
			}
			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				fmt.Println(err)
			}
		case <-ticker.C:
			//fmt.Println("Ping")
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte("")); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
func (c *Client) ReadMessage() { // from client to server
	defer func() {
		c.manager.removeClient(c, c.room.name)

	}()
	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		fmt.Println(err)
		return
	}
	c.connection.SetReadLimit(512)
	c.connection.SetPongHandler(c.pongHandler)
	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(err)
			}
			break
		}
		var request UserEvent
		request.from = c.name
		if err := json.Unmarshal(payload, &request.event); err != nil {
			log.Println(err)
			break

		}

		if err := c.manager.routeEvents(request, c); err != nil {
			log.Println(err)

		}
	}
}
func (c *Client) pongHandler(msg string) error {
	//log.Println("Pong")
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}
