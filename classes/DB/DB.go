package DB

import (
	"database/sql"

	"github.com/Smadarl/wwf/classes/Config"

	//For ease of use with sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var (
	connection *sql.DB
	err        error
)

//GetConnection - Gets persistent db connection
func GetConnection() (*sql.DB, error) {
	if connection != nil {
		err = connection.Ping()
		if err == nil {
			return connection, nil
		}
	}
	dbConfig := Config.GetDatabase()
	connection, err = sql.Open("sqlite3", dbConfig.Path)
	if err != nil {
		err = errors.Wrap(err, "Open DB")
	}

	return connection, err
}
