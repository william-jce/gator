package main

import (
	"database/sql"

	"github.com/william-jce/gator/internal/config"
	"github.com/william-jce/gator/internal/database"
)

type state struct {
	db    *database.Queries
	sqlDB *sql.DB
	cfg   *config.Config
}
