package main

import (
	"fmt"
	"net/http"

	"github.com/lafetz/trivia-go/internal/socketComm"
)

func (app *application) createGame(w http.ResponseWriter, r *http.Request) {
	name := "xof"
	room := socketComm.NewRoom(name)
	conn, err := socketComm.WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	app.manager.AddRoom(room, name)
	client := socketComm.NewClient(conn, app.manager, room)
	app.manager.JoinRoom(client, name)
	go client.ReadMessage()
	go client.SendMessage()
}
func (app *application) joinGame(w http.ResponseWriter, r *http.Request) {
	name := "xof"
	room := app.manager.GetRoom(name)

	conn, err := socketComm.WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	client := socketComm.NewClient(conn, app.manager, room)
	app.manager.JoinRoom(client, name)
	go client.ReadMessage()
	go client.SendMessage()
}
