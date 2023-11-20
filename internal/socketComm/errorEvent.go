package socketComm

import "fmt"

type Error struct {
	message string
	code    int32
}

func newError(message string, code int32) Error {
	return Error{
		message: message,
		code:    code,
	}
}
func errorEvent(message string, code int32, c *Client) {
	x := newError(message, code)
	event, err := createEvent(EventError, x)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.egress <- event
}
