package main

import (
	"log"
	"net/http"

	"github.com/HLJman/AptCoPilot/internal/storage"
	"github.com/HLJman/AptCoPilot/route"
	goji "goji.io"
)

func main() {
	err := storage.Connect()
	if err != nil {
		log.Fatal(err)
	}

	mux := goji.NewMux()
	route.Register(mux)
	go func() {
		err := http.ListenAndServe(":8081", mux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = http.ListenAndServe(":8080", http.FileServer(http.Dir("./assets")))
	if err != nil {
		log.Fatal(err)
	}
}
