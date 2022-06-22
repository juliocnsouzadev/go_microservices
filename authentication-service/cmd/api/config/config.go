package config

import (
	"authetication/data"
	"database/sql"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}
