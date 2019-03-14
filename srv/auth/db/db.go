package db

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/zebresel-com/mongodm"
)

var connection *mongodm.Connection

// D returns pointer to db connection
func D() *mongodm.Model {
	return connection.Model("User")
}

// Close close ODM
func Close() {
	if connection != nil {
		connection.Close()
	}
}

// Connect creates connection
func Connect() (*mongodm.Connection, error) {

	// Load error texts
	var localMap map[string]map[string]string
	file, err := ioutil.ReadFile("./db/locals.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(file, &localMap); err != nil {
		return nil, err
	}

	mongoc := &mongodm.Config{
		DatabaseHosts:    []string{os.Getenv("DB_HOST")},
		DatabaseName:     os.Getenv("DB_NAME"),
		DatabaseUser:     os.Getenv("DB_USER"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
		DatabaseSource:   "admin",
		Locals:           localMap["en-US"],
	}

	cnn, err := mongodm.Connect(mongoc)
	if err != nil {
		return nil, err
	}

	connection = cnn
	return cnn, nil
}
