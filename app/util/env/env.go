package env

import (
	"os"
	"strings"
)

func DBDriver() string {
	return os.Getenv("DB_DRIVER")
}

func DBUser() string {
	return os.Getenv("DB_USER")
}

func DBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func DBHost() string {
	return os.Getenv("DB_HOST")
}

func DBName() string {
	return os.Getenv("DB_NAME")
}

func DBPort() string {
	return os.Getenv("DB_PORT")
}

func ServerPortNo() string {
	return os.Getenv("SERVER_PORT_NO")
}

func AllowOrigins() []string {
	return strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
}

func DisableAuth() bool {
	return os.Getenv("DISABLE_AUTH") != ""
}

func BuildVersion() string {
	return os.Getenv("BUILD_VERSION")
}

func AppID() string {
	return os.Getenv("APP_ID")
}
