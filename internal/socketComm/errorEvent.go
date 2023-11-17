package socketComm

type Error struct {
	message string
	code    int32
}

func createErrorResponse(message string, code int32) Error {
	return Error{
		message: message,
		code:    code,
	}
}
func errorEvent(message string, code int32, c *Client) {
	x := createErrorResponse(message, code)
	event := createEvent(EventError, x)
	c.egress <- event

}
