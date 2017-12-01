package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/HLJman/AptCoPilot/internal/storage"
	"github.com/HLJman/AptCoPilot/route"
	goji "goji.io"
)

var (
	username string
	password string
	server   string
	database string
)

func init() {
	flag.StringVar(&username, "dbusername", "jadmin", "username for database")
	flag.StringVar(&password, "dbpassword", "a00787206", "password for database")
	flag.StringVar(&server, "dbserver", "localhost:8889", "server url for database")
	flag.StringVar(&database, "dbname", "copilot", "name for database")

	flag.Parse()
}

func main() {
	err := storage.Connect(username, password, server, database)
	if err != nil {
		log.Fatal(err)
	}

	mux := goji.NewMux()
	route.Register(mux)
	go func() {
		err := http.ListenAndServe(":8001", mux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = http.ListenAndServe(":8000", http.FileServer(http.Dir("./assets")))
	if err != nil {
		log.Fatal(err)
	}
}
