package db

import (
	"database/sql"
	"log"
	"time"
	"os"
	
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute


type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// creates database pool for Postgres
func ConnectSQL() (*DB, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// db, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URI"))
	d, err := NewDatabase(os.Getenv("POSTGRES_URI"))
	
	if err != nil {
		log.Fatalln(err.Error())
	}
	
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	err = TestDB(d)

	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// tries to ping the db (check conn)
func TestDB(d *sql.DB) error {
	err := d.Ping()

	if err != nil {
		return err
	}

	return nil
}

// creates a new db for the app
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}