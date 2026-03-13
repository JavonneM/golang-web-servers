package database

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/sensitive"
	_ "github.com/lib/pq" // Driver import
)

type DatabaseWrapper interface {
	Close() error
}

type DatabaseImpl struct {
	Conn *sql.DB
}

func NewPostgresDatabase(host string, databasename string, username sensitive.String, password sensitive.String, port string) (DatabaseWrapper, error) {
	connection, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"%s:%s@%s:%s/%s",
			string(username),
			string(password),
			host,
			port,
			databasename,
		),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("database connection established")
	return DatabaseImpl{
		Conn: connection,
	}, nil
}

func (db DatabaseImpl) Close() error {
	return db.Conn.Close()
}
