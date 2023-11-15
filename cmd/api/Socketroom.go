package main

type RoomList map[string]Room
type Room struct {
	clientList ClientList
	owner      *Client
	name       string
}
