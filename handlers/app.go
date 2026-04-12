package handlers

import (
	"database/sql"
	"github.com/chiprek/bassurance/store"
)

type App struct {
	DB *sql.DB
}
