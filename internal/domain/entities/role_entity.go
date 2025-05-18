package entities

import (
	"football_licence/shared/app_const"
	"gorm.io/gorm"
	"time"
)

type RoleEntity struct {
	Id        uint           `gorm:"primaryKey"`
	TitleRu   string         `gorm:"type:varchar(255);not null;comment:Название на русском" json:"title_ru" form:"title_ru" validate:"required"`
	TitleKk   string         `gorm:"type:varchar(255);not null;comment:Атауы қазақша" json:"title_kk" form:"title_kk" validate:"required"`
	TitleEn   *string        `gorm:"type:varchar(255);default:NULL;comment:Title in English" json:"title_en" form:"title_en" validate:"omitempty,max=255"`
	Value     string         `gorm:"type:varchar(300);not null;uniqueIndex;comment:Уникальное значение" json:"value" form:"value" validate:"required"`
	IsActive  bool           `gorm:"default:true;not null; comment:Активна ли роль" json:"is_active" form:"is_active" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (RoleEntity) TableName() string {
	return app_const.RoleTableName
}
