package role_case

import (
	"football_licence/internal/domain/entities"
	"football_licence/internal/dto/role_dto"
	"football_licence/internal/repository/role_repository"
	"github.com/dranikpg/dto-mapper"
	"github.com/go-playground/validator/v10"
)

type CreateRoleUseCase struct {
	repo      *role_repository.RoleRepo
	validator *validator.Validate
}

func NewCreateRoleUseCase(repo *role_repository.RoleRepo) *CreateRoleUseCase {
	return &CreateRoleUseCase{
		repo:      repo,
		validator: validator.New(),
	}
}

func (uc *CreateRoleUseCase) Validate(input role_dto.RoleCreateDTO) error {
	return uc.validator.Struct(input)
}

func (uc *CreateRoleUseCase) Transform(input role_dto.RoleCreateDTO) *entities.RoleEntity {
	var entity entities.RoleEntity
	dto.Map(&entity, input)
	return &entity
}

func (uc *CreateRoleUseCase) Execute(input role_dto.RoleCreateDTO) (*entities.RoleEntity, error) {
	if err := uc.Validate(input); err != nil {
		return nil, err
	}
	entity := uc.Transform(input)
	if err := uc.repo.Create(entity); err != nil {
		return nil, err
	}
	return entity, nil
}
