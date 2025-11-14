package responses

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	database "github.com/spiritov/jump/api/db/queries"
)

var (
	Queries *database.Queries
)

func OpenDB(path string) *sql.DB {
	// os.Remove("./db/jump.db")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("[fatal] failed to open db: %v", err)
	}

	Queries = database.New(db)
	return db
}
