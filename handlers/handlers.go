package handlers

import (
	"database/sql"

	"github.com/sidmohanty11/gradbook/server/db"
)

// creates a global postgresdb instance.
var Psql *sql.DB

//creates a new repository, gets the db conn from the main.go file
func NewRepo(db *db.DB) {
	Psql = db.SQL
}
