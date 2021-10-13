package config

import (
	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Settings struct {
	Host string
	Port int
	Addr string

	DatabaseURL  string
	DatabaseName string
}

var currentSettings *Settings

func envStr(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultValue
}

func envInt(key string, defaultValue int) int {
	if val, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return val
	}

	return defaultValue
}

func Load() {
	host := envStr("ADDRESS", "")
	port := envInt("PORT", 5000)

	dbURL := envStr("DATABASE_NAME", "mongodb://localhost")
	dbName := "dockwork"
	if dbConn, err := connstring.ParseAndValidate(dbURL); err == nil && dbConn.Database != "" {
		dbName = dbConn.Database
	}

	currentSettings = &Settings{
		Host: host,
		Port: port,
		Addr: fmt.Sprintf("%s:%d", host, port),

		DatabaseURL:  envStr("DATABASE_URL", dbURL),
		DatabaseName: envStr("DATABASE_NAME", dbName),
	}
}

func Get() *Settings {
	return currentSettings
}
