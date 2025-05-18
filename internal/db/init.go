package db

import (
	"football_licence/config"
)

func RegisterDB(cfg *config.AppConfig) {
	InitDatabase(cfg)
	MigrateDB()
}
