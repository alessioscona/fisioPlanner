package app

import (
	"fisioplanner/internal/db"
	"fmt"
)

func Run() {
    database, err := db.Connect()
    if err != nil {
        panic(err)
    }
    defer database.Close()

    err = database.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Benvenuto in FisioPlanner! Connessione al database riuscita.")
}