package role_dto

import "time"

// swagger:model RoleResponseDTO
type RoleResponseDTO struct {
	Id        uint      `json:"id"`
	Value     string    `json:"value"`
	TitleRu   string    `json:"title_ru"`
	TitleKk   string    `json:"title_kk"`
	TitleEn   *string   `json:"title_en"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
