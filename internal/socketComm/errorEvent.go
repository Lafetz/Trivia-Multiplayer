package socketComm

import "fmt"

type Error struct {
	Message string
	//code    int32
}

func newError(message string) Error {
	return Error{
		Message: message,
		//code:    code,
	}
}
func errorEvent(message string, c *Client) {
	errorS := newError(message)
	event, err := createEvent(EventError, errorS)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.egress <- event

}
func failedMessageClient(c *Client) {
	message := "failed to send messsage"
	errorEvent(message, c)
}
func serverError(err error, c *Client) {
	fmt.Println(err, "12")
	message := "the server encountered a problem and could not process your request"
	errorEvent(message, c)
}

func jsonError(err error, c *Client) {
	message := "body contains badly-formed JSON"
	errorEvent(message, c)
}
