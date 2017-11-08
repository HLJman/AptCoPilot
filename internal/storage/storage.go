package storage

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

var db *sql.DB

// $servername = "localhost:8889";
// $username = "jadmin";
// $password = "a00787206";
// $database = "copilot";

func Connect() (err error) {
	db, err = sql.Open("mysql", "jadmin:a00787206@tcp(localhost:8889)/copilot")
	if err != nil {
		return err
	}

	return nil
}

type Property struct {
	ID        string `sql:"id"`
	Name      string `sql:"name"`
	Address   string `sql:"address"`
	Latitude  string `sql:"lat"`
	Longitude string `sql:"lng"`
	Type      string `sql:"type"`
	City      string `sql:"city"`
	Units     string `sql:"units"`
	County    string `sql:"county"`
	MainPic   string `sql:"mainpic"`
}

type Properties []Property

func AllProperties() (Properties, error) {
	rows, err := db.Query("SELECT * FROM properties")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	ps := make(Properties, 0)
	for rows.Next() {
		var p Property
		if err := rows.Scan(&p); err != nil {
			log.Println(err)
		}

		ps = append(ps, p)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return ps, nil
}
