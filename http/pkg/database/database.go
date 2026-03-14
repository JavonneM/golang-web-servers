package database

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/sensitive"
	_ "github.com/lib/pq" // Driver import
)

func NewPostgresDatabase(host string, databasename string, username sensitive.String, password sensitive.String, port string) (*sql.DB, error) {
	dsname := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		string(username),
		string(password),
		databasename,
	)
	fmt.Println(dsname)

	connection, err := sql.Open(
		"postgres",
		dsname,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("database connection established")
	fmt.Println("database: ping test")
	err = connection.Ping()
	if err != nil {
		return nil, fmt.Errorf("database ping test failed %w", err)
	}
	return connection, nil
}
