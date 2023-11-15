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
	clientList ClientList
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
	fmt.Println(event.Type, string(event.Payload))
	return nil
}
func NewManager() *Manager {
	m := &Manager{
		handlers:   make(map[string]EventHandler),
		clientList: make(ClientList),
	}
	m.setupEventHandlers()
	return m
}
func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clientList[client] = true
}
func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.clientList[client]; ok {
		client.connection.Close()
		delete(m.clientList, client)
	}
}
