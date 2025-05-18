package role_repository

import (
	"football_licence/internal/domain/entities"
	"football_licence/internal/repository"
	"gorm.io/gorm"
)

type RoleRepo struct {
	*repository.BaseRepository[entities.RoleEntity]
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		BaseRepository: repository.NewBaseRepository[entities.RoleEntity](db),
	}
}
