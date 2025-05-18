package db

import (
	"database/sql"
	"fmt"
	"football_licence/config"
	"football_licence/internal/domain/entities"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase(cfg *config.AppConfig) {
	log.Println("⚙️ Initializing database...")
	if cfg.DbType == "postgresql" {
		createPostgresDatabaseIfNotExists(cfg)
	}

	var dialector gorm.Dialector

	dialector = getDBTypeConnection(cfg)

	var err error
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Произошла ошибка при подключении к  %s БД: %v", cfg.DbType, err)
	}
}

func getDBTypeConnection(cfg *config.AppConfig) gorm.Dialector {
	switch cfg.DbType {
	case "postgresql":
		if cfg.DbPgUser == nil || cfg.DbPgPassword == nil || cfg.DbPgPort == nil {
			log.Fatal("PostgreSQL credentials are missing in config")
		}
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
			cfg.DbPgHost, *cfg.DbPgUser, *cfg.DbPgPassword, cfg.DbName, *cfg.DbPgPort,
		)
		return postgres.Open(dsn)

	case "mysql":
		if cfg.DbMysqlUser == nil || cfg.DbMysqlPassword == nil || cfg.DbMysqlPort == nil {
			log.Fatal("MySQL credentials are missing in config")
		}
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			*cfg.DbMysqlUser, *cfg.DbMysqlPassword, cfg.DbMysqlHost, *cfg.DbMysqlPort, cfg.DbName,
		)
		return mysql.Open(dsn)

	default:
		log.Fatalf("Unsupported DB_TYPE: %s", cfg.DbType)
		return nil
	}
}

func createPostgresDatabaseIfNotExists(cfg *config.AppConfig) {
	if cfg.DbPgUser == nil || cfg.DbPgPassword == nil || cfg.DbPgPort == nil {
		log.Fatal("PostgreSQL credentials are missing in config")
	}

	// Подключение к системной базе postgres
	postgresDSN := fmt.Sprintf(
		"host=%s user=%s password=%s port=%d dbname=postgres sslmode=disable",
		cfg.DbPgHost, *cfg.DbPgUser, *cfg.DbPgPassword, *cfg.DbPgPort,
	)

	sqlDB, err := sql.Open("postgres", postgresDSN)
	if err != nil {
		log.Fatalf("❌ Cannot connect to system PostgreSQL DB: %v", err)
	}
	defer sqlDB.Close()

	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)"
	if err := sqlDB.QueryRow(checkQuery, cfg.DbName).Scan(&exists); err != nil {
		log.Fatalf("❌ Failed to check if database exists: %v", err)
	}

	if !exists {
		createQuery := fmt.Sprintf("CREATE DATABASE \"%s\"", cfg.DbName)
		_, err := sqlDB.Exec(createQuery)
		if err != nil {
			log.Fatalf("❌ Failed to create database %s: %v", cfg.DbName, err)
		}
		log.Printf("✅ Database %s created!", cfg.DbName)
	} else {
		log.Printf("ℹ️ Database %s already exists", cfg.DbName)
	}
}

func MigrateDB() {
	DB.AutoMigrate(&entities.RoleEntity{})
}
