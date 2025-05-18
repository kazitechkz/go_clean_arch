package config

import (
	"football_licence/shared/app_const"
	"football_licence/shared/utils"
	"github.com/joho/godotenv"
	"log"
)

type AppConfig struct {
	AppName          string
	AppBodyLimits    int
	AppCaseSensitive bool
	AppConcurrency   int
	AppStatus        string
	AppPort          int
	AppHost          string
	AppLocale        string
	DbType           string
	DbName           string
	DbPgHost         string
	DbPgUser         *string
	DbPgPassword     *string
	DbPgPort         *int
	DbMysqlHost      string
	DbMysqlUser      *string
	DbMysqlPassword  *string
	DbMysqlPort      *int
	ApiBaseUrl       string
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	return &AppConfig{
		AppName:          utils.GetString(app_const.EnvAppName),
		AppBodyLimits:    utils.GetInt(app_const.EnvAppBodyLimits),
		AppCaseSensitive: utils.GetBool(app_const.EnvAppCaseSensitive),
		AppConcurrency:   utils.GetInt(app_const.EnvAppConcurrency),
		AppStatus:        utils.GetString(app_const.EnvAppStatus),
		AppPort:          utils.GetInt(app_const.EnvAppPort),
		AppHost:          utils.GetString(app_const.EnvAppHost),
		AppLocale:        utils.GetString(app_const.EnvAppLocale),
		DbType:           utils.GetString(app_const.EnvDbType),
		DbName:           utils.GetString(app_const.EnvDbName),
		DbPgHost:         utils.GetString(app_const.EnvDbPgHost),
		DbPgUser:         utils.GetStringPtr(app_const.EnvDbPgUser),
		DbPgPassword:     utils.GetStringPtr(app_const.EnvDbPgPassword),
		DbPgPort:         utils.GetIntPtr(app_const.EnvDbPgPort),
		DbMysqlHost:      utils.GetString(app_const.EnvDbMysqlHost),
		DbMysqlUser:      utils.GetStringPtr(app_const.EnvDbMysqlUser),
		DbMysqlPassword:  utils.GetStringPtr(app_const.EnvDbMysqlPassword),
		DbMysqlPort:      utils.GetIntPtr(app_const.EnvDbMysqlPort),
		ApiBaseUrl:       utils.GetString(app_const.ApiBaseUrl),
	}
}
