package http

import (
	"football_licence/config"
	"football_licence/internal/delivery/http/role_http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func RegisterAllRoutes(app *fiber.App, db *gorm.DB, cfg *config.AppConfig) {
	// Swagger docs
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group(cfg.ApiBaseUrl)
	role_http.RegisterRoleRoutes(api, db)
}
