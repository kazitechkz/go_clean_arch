package role_http

import (
	"football_licence/shared/app_const"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoleRoutes(api fiber.Router, db *gorm.DB) {
	role_route_api := api.Group(app_const.RolesPathName)
	role_route_api.Post(app_const.CreatePathName, CreateRoleHandler(db))
	role_route_api.Get(app_const.GetByIdPathName, GetRoleByIDHandler(db))
	role_route_api.Get(app_const.GetByValuePathName, GetRoleByValueHandler(db))
}
