package main

type RoomList map[string]Room
type Room struct {
	clientList ClientList
	owner      *Client
	name       string
}

// func (room *Room) joinRoom(client *Client) {
// 	room.Lock()
// 	defer room.Unlock()
// 	room.clientList[client] = true
// }
// func (room *Room) leaveRoom(client *Client) {
// 	room.Lock()
// 	defer room.Unlock()
// 	if _, ok := room.clientList[client]; ok {
// 		client.connection.Close()
// 		delete(room.clientList, client)
// 	}
// }
