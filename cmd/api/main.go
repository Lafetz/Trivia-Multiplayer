package main

import (
	"fmt"
	"net/http"

	"github.com/lafetz/trivia-go/internal/socketComm"
)

type application struct {
	manager *socketComm.Manager
}

func main() {
	manager := socketComm.NewManager()
	app := application{
		manager: manager,
	}

	fmt.Println("server on 4000")
	err := http.ListenAndServe(":4000", app.routes())
	if err != nil {
		fmt.Println(err)
		return
	}

}
