package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HLJman/AptCoPilot/config"
	"github.com/HLJman/AptCoPilot/internal/storage"
	"github.com/HLJman/AptCoPilot/route"
	goji "goji.io"
)

func main() {
	err := storage.Connect(config.DBUsername(), config.DBPassword(), config.DBServer(), config.DBDatabase())
	if err != nil {
		log.Fatal(err)
	}

	mux := goji.NewMux()
	route.Register(mux)
	go func() {
		fmt.Println("Starting server")
		err := http.ListenAndServe(":8001", mux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = http.ListenAndServe(":80", http.FileServer(http.Dir("./assets")))
	if err != nil {
		log.Fatal(err)
	}
}
