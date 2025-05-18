package role_dto

// swagger:model RoleCreateDTO
type RoleCreateDTO struct {
	TitleRu  string  `json:"title_ru" validate:"required,max=255"`
	TitleKk  string  `json:"title_kk" validate:"required,max=255"`
	TitleEn  *string `json:"title_en" validate:"omitempty,max=255"`
	Value    string  `json:"value" validate:"required,max=300"`
	IsActive bool    `json:"is_active"`
}
