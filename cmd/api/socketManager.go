package main

import (
	"errors"
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

func NewManager() *Manager {
	m := &Manager{
		handlers: make(map[string]EventHandler),
		roomList: make(RoomList),
	}
	m.setupEventHandlers()
	return m
}
func (m *Manager) setupEventHandlers() { // functions based on events
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