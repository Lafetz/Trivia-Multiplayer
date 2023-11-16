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
	room := Room{
		clientList: make(ClientList),
		owner:      "abel",
		name:       name,
	}
	client := NewClient(conn, app.manager, &room)

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
	room := app.manager.getRoom(name)
	client := NewClient(conn, app.manager, room)

	app.manager.joinRoom(client, name)
	go client.readMessage()
	go client.sendMessage()
}
