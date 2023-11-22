package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/lafetz/trivia-go/internal/data"
	"github.com/lafetz/trivia-go/internal/socketComm"
)

type application struct {
	manager *socketComm.Manager
	db      *sql.DB
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	dsnDatabase := os.Getenv("DB_DSN")

	manager := socketComm.NewManager()
	app := application{
		manager: manager,
	}
	db, err := data.InitDB(dsnDatabase)
	if err != nil {
		fmt.Println(err)
		return
	}
	app.db = db
	fmt.Println("server on 4000")
	err = http.ListenAndServe(":4000", app.routes())
	if err != nil {
		fmt.Println(err)
		return
	}

}
