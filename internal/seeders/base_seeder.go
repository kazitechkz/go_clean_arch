package seeders

import (
	"fmt"
	"football_licence/config"
	"football_licence/internal/db"
	"football_licence/internal/seeders/role_seeder"
)

func Seeder(app_config *config.AppConfig) {
	fmt.Println("🚀 Запускаем сидеры для:", app_config.AppStatus)

	// Список всех сидеров
	seederList := []SeederInterface{
		role_seeder.RoleSeeder{},
	}

	// Запускаем каждый сидер
	for _, seeder := range seederList {
		seeder.Seed(db.DB, app_config)
	}

	fmt.Println("✅ Все сидеры выполнены!")

}
