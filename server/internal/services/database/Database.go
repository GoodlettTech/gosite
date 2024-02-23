package Database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once
var database *sql.DB

func GetInstance() *sql.DB {
	if database == nil {
		once.Do(func() {
			var err error
			database, err = sql.Open("postgres", getConnectionString())
			if err != nil {
				panic("couldn't open database")
			}

			initUsersTable(database)
		})
	}
	return database
}

func getConnectionString() string {
	var (
		host     = os.Getenv("POSTGRES_HOSTNAME")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func initUsersTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id			SERIAL PRIMARY KEY,
		email		varchar(32) UNIQUE,
		username	varchar(32) UNIQUE,
		password	varchar(32)
	);`)

	if err != nil {
		panic(err)
	}
}
