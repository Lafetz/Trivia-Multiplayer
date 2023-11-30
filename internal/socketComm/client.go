package socketComm

import (
	"encoding/json"
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
	egress     chan *Event
	name       string
}

func NewClient(conn *websocket.Conn, manager *Manager, room *Room) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		room:       room,
		egress:     make(chan *Event),
	}
}
func (c *Client) SendMessage() { //from server to client
	defer func() {
		c.manager.removeClient(c, c.room.Name)
	}()
	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return

			}
			data, err := json.Marshal(message)
			if err != nil { ///
				serverError(err, c)
			}
			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}
		case <-ticker.C:

			if err := c.connection.WriteMessage(websocket.PingMessage, []byte("")); err != nil {
				log.Println("writemsg: ", err)
				return
			}
		}
	}
}
func (c *Client) ReadMessage() { // from client to server

	defer func() {
		c.manager.removeClient(c, c.room.Name)

	}()
	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		return
	}
	c.connection.SetReadLimit(512)
	c.connection.SetPongHandler(c.pongHandler)
	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("connection closed: ", err)
			}
			break
		}
		var request UserEvent
		request.from = c.name
		if err := json.Unmarshal(payload, &request.event); err != nil {
			jsonError(err, c) //event type not**
			continue

		}

		if err := c.manager.routeEvents(request, c); err != nil {
			serverError(err, c)

		}
	}
}
func (c *Client) pongHandler(msg string) error {

	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}
