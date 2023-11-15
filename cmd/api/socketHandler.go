package main

import (
	"fmt"
	"net/http"
)

func (app *application) socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	client := NewClient(conn, app.manager)
	app.manager.addClient(client)
	go client.readMessage()
	go client.sendMessage()
}
