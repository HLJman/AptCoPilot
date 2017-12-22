package config

import (
	"io/ioutil"
	"os"
)

func DBUsername() string {
	return os.Getenv("DB_USERNAME")
}

func DBPassword() string {
	data, err := readFile(os.Getenv("DB_PASSWORD"))
	if err != nil {
		return ""
		// log.Fatal(err)
	}

	return data
}

func DBServer() string {
	return os.Getenv("DB_SERVER")
}

func DBDatabase() string {
	return os.Getenv("DB_NAME")
}

func readFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
