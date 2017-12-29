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
	if config.Local() == false {
		err := storage.Connect(config.DBUsername(), config.DBPassword(), config.DBServer(), config.DBDatabase())
		if err != nil {
			log.Fatal(err)
		}
	}

	mux := goji.NewMux()
	route.Register(mux)

	fmt.Println("Starting server")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
