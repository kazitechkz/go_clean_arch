package role_http

import (
	"errors"
	"football_licence/internal/dto/role_dto"
	"football_licence/internal/repository/role_repository"
	"football_licence/internal/use_case/role_case"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

// CreateRoleHandler godoc
// @Summary Create a new role_case
// @Tags roles
// @Accept json
// @Produce json
// @Param data body role_dto.RoleCreateDTO true "Role Data"
// @Success 201 {object} entities.RoleEntity
// @Router /api/role/create [post]
func CreateRoleHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input role_dto.RoleCreateDTO
		useCase := role_case.NewCreateRoleUseCase(role_repository.NewRoleRepo(db))
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Невалидное тело запроса"})
		}
		entity, err := useCase.Execute(input)
		if err != nil {
			return c.Status(422).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(201).JSON(entity)
	}
}

// GetRoleByIDHandler godoc
// @Summary      Получить роль по ID
// @Description  Возвращает одну роль по её идентификатору
// @Tags         roles
// @Produce      json
// @Param        id path int true "ID роли"
// @Success      200 {object} role_dto.RoleResponseDTO
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /api/role/get/{id} [get]
func GetRoleByIDHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		idUint, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректный ID"})
		}

		useCase := role_case.NewGetRoleByIDUseCase(role_repository.NewRoleRepo(db))
		role, err := useCase.Execute(uint(idUint))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Роль не найдена"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(role)
	}
}

// GetRoleByIDHandler godoc
// @Summary      Получить роль по Value
// @Description  Возвращает одну роль по её строковому идентификатору
// @Tags         roles
// @Produce      json
// @Param        value path string true "Value роли"
// @Success      200 {object} role_dto.RoleResponseDTO
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /api/role/get-by-value/{value} [get]
func GetRoleByValueHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		value := c.Params("value")
		useCase := role_case.NewGetRoleByValueUseCase((role_repository.NewRoleRepo(db)))
		role, err := useCase.Execute(value)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Роль не найдена"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(role)
	}
}
