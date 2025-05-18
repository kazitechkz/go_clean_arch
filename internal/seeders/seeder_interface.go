package seeders

import (
	"football_licence/config"
	"gorm.io/gorm"
)

type SeederInterface interface {
	PrepareBeforeCreate(db *gorm.DB)                // Очистка таблицы перед созданием
	GetProdData() interface{}                       // Данные для продакшена
	GetDevData() interface{}                        // Данные для разработки
	Seed(db *gorm.DB, app_config *config.AppConfig) // Запуск сидера
}
