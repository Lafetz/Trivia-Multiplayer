package socketComm

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	WebsocketUpgrader = websocket.Upgrader{
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
	m.handlers[EventSendAnswer] = SendAnswr
	m.handlers[EventGameStart] = StartGame
}
func (m *Manager) routeEvents(UserEvent UserEvent, c *Client) error {
	if handler, ok := m.handlers[UserEvent.event.Type]; ok {
		if err := handler(UserEvent.event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("no such event type")
	}

}

func (m *Manager) JoinRoom(client *Client, name string) {
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
func (m *Manager) AddRoom(room *Room, name string) {
	m.Lock()
	defer m.Unlock()
	m.roomList[name] = room
}
func (m *Manager) GetRoom(name string) *Room {
	room := m.roomList[name]
	return room
}
