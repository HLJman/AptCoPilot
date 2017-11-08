package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HLJman/AptCoPilot/internal/storage"
	"goji.io/pat"
)

func Properties(w http.ResponseWriter, r *http.Request) {
	var ps storage.Properties
	var err error

	switch pat.Param(r, "type") {
	case "all":
		ps, err = storage.AllProperties()
		if err != nil {
			log.Println(err)
			return
		}
	case "marketrate":
		ps, err = storage.MarketRateProperties()
		if err != nil {
			log.Println(err)
			return
		}
	case "affordable":
		ps, err = storage.AffordableProperties()
		if err != nil {
			log.Println(err)
			return
		}
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
}
