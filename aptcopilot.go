package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HLJman/AptCoPilot/config"
	"github.com/HLJman/AptCoPilot/internal/storage"
	"github.com/HLJman/AptCoPilot/route"
	goji "goji.io"
	"goji.io/pat"
)

func main() {
	if config.Local() {
		err := storage.Connect(config.DBUsername(), config.DBPassword(), config.DBServer(), config.DBDatabase())
		if err != nil {
			log.Fatal(err)
		}
	}

	mux := goji.NewMux()
	route.Register(mux)
	mux.Handle(pat.New("/src"), http.FileServer(http.Dir("./assets/src")))
	mux.HandleFunc(pat.New("/properties"), func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/index.html")
	})
	mux.HandleFunc(pat.New("/contact"), func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/index.html")
	})
	mux.Handle(pat.New("/*"), http.FileServer(http.Dir("./assets")))

	fmt.Println("Starting server")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
