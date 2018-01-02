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

	switch r.URL.Query().Get("type") {
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
	default:
		ps, err = storage.AllProperties()
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

func Property(w http.ResponseWriter, r *http.Request) {
	type rating struct {
		Name   string
		Rating string
	}

	type owner struct {
		Name  interface{}
		Phone string
		Email string
		Notes string
	}

	type property struct {
		ID       string
		Name     string
		Location interface{}
		Units    interface{}
		Year     string
		Images   []string
		Ratings  []rating
		Owner    owner
	}

	type unit struct {
		Units     int
		BB        string
		Size      int
		Rent      string
		RentPerSq string
	}

	json.NewEncoder(w).Encode(property{
		ID:   pat.Param(r, "id"),
		Name: "Heritage Magna",
		Year: "1980",
		Location: struct {
			Latitude  string
			Longitude string
			Address   string
			City      string
			State     string
			Zip       string
			County    string
		}{
			Latitude:  "40.69519",
			Longitude: "-112.086815",
			Address:   "15222 S Dogwood Ln",
			City:      "Lehi",
			State:     "UT",
			Zip:       "84065",
			County:    "Utah",
		},
		Units: struct {
			Total     int
			Breakdown []unit
			Vacant    unit
		}{
			Total: 100,
			Breakdown: []unit{
				unit{
					Units:     57,
					BB:        "2/1",
					Size:      880,
					Rent:      "$1200",
					RentPerSq: "$1.36",
				},
				unit{
					Units:     43,
					BB:        "2/2",
					Size:      950,
					Rent:      "$1250",
					RentPerSq: "$1.31",
				},
			},
			Vacant: unit{
				Units:     2,
				Size:      880,
				Rent:      "$1200",
				RentPerSq: "$1.31",
			},
		},
		Images: []string{
			"https://camo.mybb.com/e01de90be6012adc1b1701dba899491a9348ae79/687474703a2f2f7777772e6a71756572797363726970742e6e65742f696d616765732f53696d706c6573742d526573706f6e736976652d6a51756572792d496d6167652d4c69676874626f782d506c7567696e2d73696d706c652d6c69676874626f782e6a7067",
			"https://www.w3schools.com/w3css/img_fjords.jpg",
			"https://www.w3schools.com/w3css/img_lights.jpg",
		},
		Ratings: []rating{
			rating{
				Name:   "Location",
				Rating: "A+",
			},
			rating{
				Name:   "Property Condition",
				Rating: "B",
			},
			rating{
				Name:   "Amenities",
				Rating: "B",
			},
			rating{
				Name:   "Access to Service",
				Rating: "A-",
			},
		},
		Owner: owner{
			Name: struct {
				First string
				Last  string
			}{
				First: "Josh",
				Last:  "Bell",
			},
			Email: "happy@dog.com",
			Notes: "email preference",
			Phone: "801-555-7829",
		},
	})
}
