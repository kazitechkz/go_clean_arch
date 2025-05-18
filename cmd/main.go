package main

import (
	"fmt"
	"football_licence/config"
	_ "football_licence/docs"
	"football_licence/internal/db"
	"football_licence/internal/delivery/http"
	"football_licence/internal/seeders"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: cfg.AppCaseSensitive,
		StrictRouting: true,
		ServerHeader:  cfg.AppName + " API",
		AppName:       cfg.AppName,
		BodyLimit:     cfg.AppBodyLimits,
		Concurrency:   cfg.AppConcurrency,
	})
	// Инициализация БД через GORM
	db.RegisterDB(cfg)
	seeders.Seeder(cfg)
	//Инициализация роутов
	http.RegisterAllRoutes(app, db.DB, cfg)

	address := fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort)
	if err := app.Listen(address); err != nil {
		log.Fatalf("❌ Ошибка Сервера: %v", err)
	}

}
