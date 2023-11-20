package socketComm

import (
	"encoding/json"
	"strconv"
)

type UserEvent struct {
	event *Event
	from  string
}
type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
type EventHandler func(event *Event, c *Client) error

const (
	//sent from user to server
	EventSendMessage = "send_message"
	EventSendAnswer  = "send_answer"
	//from server
	EventError        = "error"
	EventGameStart    = "game_start"
	EventSendQuestion = "send_Question"
	EventFinalScores  = "final_scores"
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

func SendMessage(event *Event, c *Client) error {
	failedMessageClient(c)
	forwardPayload(event, c.room)
	return nil
}

func SendAnswr(event *Event, c *Client) error {

	if c.answer != 0 {
		// fmt.Println("user submited again")
		return nil

	}
	userAnswer, err := strconv.Atoi(string(event.Payload))

	if err != nil {

		return err
	}

	if int32(userAnswer) == c.room.CurrentQuestion.Answer {
		c.score++
	}
	c.answer = int32(userAnswer)
	forwardPayload(event, c.room)
	return nil
}
func StartGame(event *Event, c *Client) error {
	go c.room.startGame()
	return nil
}

func SendScores(room *Room) error {
	array := []UserScore{}
	for c := range room.clientList {
		array = append(array, UserScore{
			Name:  c.name,
			Score: int(c.score),
		})
	}
	event, err := createEvent(EventFinalScores, &ScoresList{Scores: array})
	if err != nil {

		return err
	}
	forwardPayload(event, room)
	return nil
}

func sendQuestion(room *Room) error {
	event, err := createEvent(EventSendQuestion, room.CurrentQuestion)
	if err != nil {
		return err
	}
	forwardPayload(event, room)
	return nil
}

// helper fucs
func createEvent(eventType string, payload interface{}) (*Event, error) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	event := Event{
		Type:    eventType,
		Payload: payloadJson,
	}
	return &event, nil
}
func forwardPayload(event *Event, room *Room) { //sends payload to everyone in a room
	for c := range room.clientList {
		c.egress <- event
	}

}
