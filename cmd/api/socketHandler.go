package main

import (
	"fmt"
	"net/http"
)

func (app *application) createGame(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	name := "xof"
	client := NewClient(conn, app.manager, name)

	room := Room{
		clientList: make(ClientList),
		owner:      client,
		name:       name,
	}
	app.manager.addRoom(&room, name)
	app.manager.joinRoom(client, name)
	go client.readMessage()
	go client.sendMessage()
}
func (app *application) joinGame(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	name := "xof"
	client := NewClient(conn, app.manager, name)

	app.manager.joinRoom(client, name)
	go client.readMessage()
	go client.sendMessage()
}
