package config

import (
	"os"
)

func HTTPAddr() string {
	addr := os.Getenv("HTTP_ADDR")
	if addr != "" {
		return addr
	}
	return ":8080"
}

func DBConnectionString() string {
	db := os.Getenv("DB")
	if db != "" {
		return db
	}
	return "postgres://clykk_atif:frdeswaq@localhost/clykk?sslmode=disable"
}

func Domain() string {
	domain := os.Getenv("DOMAIN")
	if domain != "" {
		return domain
	}
	return "localhost:8080"
}
