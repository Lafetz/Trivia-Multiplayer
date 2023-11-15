package main

import (
	"fmt"
	"net/http"
)

type application struct {
	manager *Manager
}

func main() {
	manager := NewManager()
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
