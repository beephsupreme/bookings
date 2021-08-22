package dbrepo

import (
	"database/sql"
	"github.com/beephsupreme/bookings/internal/config"
	"github.com/beephsupreme/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostGresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
