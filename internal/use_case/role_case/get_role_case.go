package role_case

import (
	"football_licence/internal/dto/role_dto"
	"football_licence/internal/repository/role_repository"
	"github.com/dranikpg/dto-mapper"
)

type GetRoleByIDUseCase struct {
	repo *role_repository.RoleRepo
}

func NewGetRoleByIDUseCase(repo *role_repository.RoleRepo) *GetRoleByIDUseCase {
	return &GetRoleByIDUseCase{repo: repo}
}

func (uc *GetRoleByIDUseCase) Execute(id uint) (*role_dto.RoleResponseDTO, error) {
	entity, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}
	var output role_dto.RoleResponseDTO
	dto.Map(&output, entity)
	return &output, nil
}
