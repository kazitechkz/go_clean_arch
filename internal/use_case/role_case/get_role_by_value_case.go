package role_case

import (
	"football_licence/internal/dto/role_dto"
	"football_licence/internal/filter"
	"football_licence/internal/repository/role_repository"
	"football_licence/shared/utils"
	"github.com/dranikpg/dto-mapper"
	"net/url"
)

type GetRoleByValueUseCase struct {
	repo *role_repository.RoleRepo
}

func NewGetRoleByValueUseCase(repo *role_repository.RoleRepo) *GetRoleByValueUseCase {
	return &GetRoleByValueUseCase{repo: repo}
}

func (uc *GetRoleByValueUseCase) Execute(value string) (*role_dto.RoleResponseDTO, error) {
	safeSearch := "%" + utils.EscapeLikePattern(url.QueryEscape(value)) + "%"
	filters := []filter.GenericFilter{
		filter.GenericFilter{
			Field: "value",
			Op:    "ILIKE",
			Value: safeSearch,
		},
	}
	entity, err := uc.repo.GetFirstWithDynamicFilters(filters, "Id", "desc")
	if err != nil {
		return nil, err
	}
	var output role_dto.RoleResponseDTO
	dto.Map(&output, entity)
	return &output, nil
}
