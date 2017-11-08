package storage

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// $servername = "localhost:8889";
// $username = "jadmin";
// $password = "a00787206";
// $database = "copilot";

func Connect() (err error) {
	db, err = sqlx.Connect("mysql", "jadmin:a00787206@tcp(localhost:8889)/copilot")
	if err != nil {
		return err
	}

	return nil
}

type Property struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	Address   string  `db:"address"`
	Latitude  float32 `db:"lat"`
	Longitude float32 `db:"lng"`
	Type      string  `db:"type"`
	City      string  `db:"city"`
	Units     int     `db:"units"`
	County    string  `db:"county"`
	MainPic   string  `db:"mainpic"`
	Year      string  `db:"year"`
	State     string  `db:"state"`
	Status    string  `db:"status"`
}

type Properties []Property

func AllProperties() (Properties, error) {
	ps := make(Properties, 0)
	err := db.Select(&ps, "SELECT * FROM properties")
	if err != nil {
		log.Println(err)
	}

	return ps, nil
}

func MarketRateProperties() (Properties, error) {
	ps := make(Properties, 0)
	err := db.Select(&ps, "SELECT * FROM PROPERTIES WHERE type like '%market%'")
	if err != nil {
		log.Println(err)
	}

	return ps, nil
}

func AffordableProperties() (Properties, error) {
	ps := make(Properties, 0)
	err := db.Select(&ps, "SELECT * FROM PROPERTIES WHERE type like '%affordable%'")
	if err != nil {
		log.Println(err)
	}

	return ps, nil
}
