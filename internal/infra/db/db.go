package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	connection *sql.DB
}

type DBCredentials struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewDB() *DB {
	return &DB{
		connection: nil,
	}
}

func (db *DB) Connect(credentials DBCredentials) error {
	dbCredentials := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		credentials.User,
		credentials.Password,
		credentials.Host,
		credentials.Port,
		credentials.Database,
	)
	conn, err := sql.Open("mysql", dbCredentials)
	if err != nil {
		return err
	}
	db.connection = conn
	return nil
}

func (db *DB) Close() error {
	if db.connection == nil {
		return fmt.Errorf("connection not established")
	}
	err := db.connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Query(query string, args ...any) (*sql.Rows, error) {
	if db.connection == nil {
		return nil, fmt.Errorf("connection not established")
	}
	return db.connection.Query(query, args...)
}
