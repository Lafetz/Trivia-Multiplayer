package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	roomList RoomList
	sync.RWMutex
	handlers map[string]EventHandler
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessage
}
func (m *Manager) routeEvents(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("no such event type")
	}

}
func SendMessage(event Event, c *Client) error {
	for k, v := range c.manager.roomList {
		if k == c.room {
			for k, _ := range v.clientList {
				k.egress <- event
			}
		}
	}
	fmt.Println(event.Type, string(event.Payload))
	return nil
}
func NewManager() *Manager {
	m := &Manager{
		handlers: make(map[string]EventHandler),
		roomList: make(RoomList),
	}
	m.setupEventHandlers()
	return m
}
func (m *Manager) joinRoom(client *Client, name string) {
	m.Lock()
	defer m.Unlock()
	m.roomList[name].clientList[client] = true
}
func (m *Manager) removeClient(client *Client, name string) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.roomList[name].clientList[client]; ok {
		client.connection.Close()
		delete(m.roomList[name].clientList, client)
	}
}
func (m *Manager) addRoom(room *Room, name string) {
	m.Lock()
	defer m.Unlock()
	m.roomList[name] = *room
}
