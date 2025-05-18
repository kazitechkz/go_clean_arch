package role_seeder

import (
	"fmt"
	"football_licence/config"
	"football_licence/internal/domain/entities"
	"football_licence/shared/app_const"
	"gorm.io/gorm"
)

type RoleSeeder struct{}

func (seeder RoleSeeder) Seed(db *gorm.DB, app_config *config.AppConfig) {
	var count int64
	db.Model(&entities.RoleEntity{}).Count(&count)
	if count > 0 {
		fmt.Println("✅ Roles уже заполнены, пропускаем...")
		return
	}

	seeder.PrepareBeforeCreate(db)

	// Определяем, какие данные загружать (Prod или Dev)
	var data interface{}
	if app_config.AppStatus == "prod" {
		data = seeder.GetProdData()
	} else {
		data = seeder.GetDevData()
	}

	db.Create(data)
	fmt.Println("✅ Roles добавлены!")
}

func (seeder RoleSeeder) PrepareBeforeCreate(db *gorm.DB) {
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", app_const.RoleTableName))
}

func (seeder RoleSeeder) GetProdData() interface{} {
	return []entities.RoleEntity{
		{
			TitleRu:  "Администратор клуба",
			TitleKk:  "Клуб әкімшісі",
			Value:    "club_admin",
			IsActive: true,
		},
		{
			TitleRu:  "Юридический специалист",
			TitleKk:  "Заңгер маман",
			Value:    "legal_specialist",
			IsActive: true,
		},
		{
			TitleRu:  "Финансовый специалист",
			TitleKk:  "Қаржы маманы",
			Value:    "finance_specialist",
			IsActive: true,
		},
		{
			TitleRu:  "Спортивный директор",
			TitleKk:  "Спорт директоры",
			Value:    "sport_director",
			IsActive: true,
		},
		{
			TitleRu:  "Администратор системы",
			TitleKk:  "Жүйе әкімшісі",
			Value:    "system_admin",
			IsActive: true,
		},
		{
			TitleRu:  "Департамент лицензирования",
			TitleKk:  "Лицензиялау департаменті",
			Value:    "licensing_department",
			IsActive: true,
		},
		{
			TitleRu:  "Юридический департамент",
			TitleKk:  "Заң департаменті",
			Value:    "legal_department",
			IsActive: true,
		},
		{
			TitleRu:  "Финансовый департамент",
			TitleKk:  "Қаржы департаменті",
			Value:    "finance_department",
			IsActive: true,
		},
		{
			TitleRu:  "Инфраструктурный отдел",
			TitleKk:  "Инфрақұрылым бөлімі",
			Value:    "infrastructure_department",
			IsActive: true,
		},
		{
			TitleRu:  "Контрольный отдел",
			TitleKk:  "Бақылау бөлімі",
			Value:    "control_department",
			IsActive: true,
		},
	}
}

func (seeder RoleSeeder) GetDevData() interface{} {
	return []entities.RoleEntity{
		{
			TitleRu:  "Администратор клуба",
			TitleKk:  "Клуб әкімшісі",
			Value:    "club_admin",
			IsActive: true,
		},
		{
			TitleRu:  "Юридический специалист",
			TitleKk:  "Заңгер маман",
			Value:    "legal_specialist",
			IsActive: true,
		},
		{
			TitleRu:  "Финансовый специалист",
			TitleKk:  "Қаржы маманы",
			Value:    "finance_specialist",
			IsActive: true,
		},
		{
			TitleRu:  "Спортивный директор",
			TitleKk:  "Спорт директоры",
			Value:    "sport_director",
			IsActive: true,
		},
		{
			TitleRu:  "Администратор системы",
			TitleKk:  "Жүйе әкімшісі",
			Value:    "system_admin",
			IsActive: true,
		},
		{
			TitleRu:  "Департамент лицензирования",
			TitleKk:  "Лицензиялау департаменті",
			Value:    "licensing_department",
			IsActive: true,
		},
		{
			TitleRu:  "Юридический департамент",
			TitleKk:  "Заң департаменті",
			Value:    "legal_department",
			IsActive: true,
		},
		{
			TitleRu:  "Финансовый департамент",
			TitleKk:  "Қаржы департаменті",
			Value:    "finance_department",
			IsActive: true,
		},
		{
			TitleRu:  "Инфраструктурный отдел",
			TitleKk:  "Инфрақұрылым бөлімі",
			Value:    "infrastructure_department",
			IsActive: true,
		},
		{
			TitleRu:  "Контрольный отдел",
			TitleKk:  "Бақылау бөлімі",
			Value:    "control_department",
			IsActive: true,
		},
	}
}
