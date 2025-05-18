package role_dto

// swagger:model RoleUpdateDTO
type RoleUpdateDTO struct {
	TitleRu  *string `json:"title_ru,omitempty" validate:"omitempty,max=255"`
	TitleKk  *string `json:"title_kk,omitempty" validate:"omitempty,max=255"`
	TitleEn  *string `json:"title_en,omitempty" validate:"omitempty,max=255"`
	Value    string  `json:"value" validate:"required,max=300"`
	IsActive *bool   `json:"is_active,omitempty"`
}
