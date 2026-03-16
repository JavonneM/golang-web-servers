package repository

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func Migrate(conn *sql.DB) error {
	fmt.Println("migration: starting")
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("%w: unable to configure migrator: %v", ErrUnableToMigrate, err)
	}
	err = goose.Up(conn, "internal/myapp/db/migrations")
	if err != nil {
		return fmt.Errorf("%w: migration failed: %v", ErrUnableToMigrate, err)
	}
	fmt.Println("migration: complete")
	return nil
}
