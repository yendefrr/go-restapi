package apiserver

import (
	"database/sql"
	"go/rest-api/internal/app/store/sqlstore"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	server := newServer(store)

	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
