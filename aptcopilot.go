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
	mux.Use(corsHandler)
	mux.HandleFunc(pat.Get("/properties/all"), func(w http.ResponseWriter, r *http.Request) {
		ps, err := storage.AllProperties()
		if err != nil {
			log.Println(err)
			return
		}

		type property struct {
			ID        int     `json:"id"`
			Name      string  `json:"name"`
			Address   string  `json:"address"`
			Latitude  float32 `json:"lat"`
			Longitude float32 `json:"lng"`
			Type      string  `json:"type"`
			City      string  `json:"city"`
			Units     int     `json:"units"`
			County    string  `json:"county"`
			MainPic   string  `json:"mainpic"`
		}

		newProps := make([]property, 0)
		for _, p := range ps {
			newProps = append(newProps, property{
				ID:        p.ID,
				Name:      p.Name,
				Address:   p.Address,
				Latitude:  p.Latitude,
				Longitude: p.Longitude,
				Type:      p.Type,
				City:      p.City,
				Units:     p.Units,
				County:    p.County,
				MainPic:   p.MainPic,
			})
		}

		err = json.NewEncoder(w).Encode(newProps)
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

func corsHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("cors")
		if r.Method == "OPTIONS" {
			log.Println("options")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "x-pingpong")
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
