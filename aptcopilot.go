package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HLJman/AptCoPilot/internal/storage"
	goji "goji.io"
	"goji.io/pat"
)

func main() {
	err := storage.Connect()
	if err != nil {
		log.Fatal(err)
	}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/properties/all"), func(w http.ResponseWriter, r *http.Request) {
		ps, err := storage.AllProperties()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(ps)

		err = json.NewEncoder(w).Encode(ps)
		if err != nil {
			log.Println(err)
		}
	})

	go http.ListenAndServe(":8080", http.FileServer(http.Dir("./assets")))

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
