package config

import (
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

var (
	APP_ENV      = os.Getenv("APP_ENV")
	APP_RESET, _ = strconv.ParseBool(os.Getenv("APP_RESET"))

	APP_PORT      = os.Getenv("APP_PORT")
	DB_NAME       = os.Getenv("DB_NAME")
	DB_PASS       = os.Getenv("DB_PASS")
	DB_USER       = os.Getenv("DB_USER")
	DB_HOST       = os.Getenv("DB_HOST")
	DB_PORT       = os.Getenv("DB_PORT")
	SQL_CONN_NAME = os.Getenv("DB_CLOUD_CONN")
	DB_SOCKET     = os.Getenv("DB_SOCKET")

	EMAIL_RESET     = os.Getenv("EMAIL_RESET")
	EMAIL_RESET_PWD = os.Getenv("EMAIL_RESET_PWD")
	EMAIL_HOST      = os.Getenv("EMAIL_HOST")
	EMAIL_PORT      = os.Getenv("EMAIL_PORT")

	URL_BASE = os.Getenv("URL_BASE")

	ALLOWED_HOSTS = strings.Split(os.Getenv("ALLOWED_HOSTS"), ",")
)
