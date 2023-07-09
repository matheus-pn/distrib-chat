package app

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	fmt.Println("[Database] Connecting to Postgres Database..")

	host := "localhost"
	user := "postgres"
	password := "postgres"
	dbname := "distribchat-development"
	port := "5432"
	sslmode := "prefer"
	timezone := "UTC"

	if env, ok := os.LookupEnv("PG_HOST"); ok {
		host = env
	}
	if env, ok := os.LookupEnv("PG_USER"); ok {
		user = env
	}
	if env, ok := os.LookupEnv("PG_PASSWORD"); ok {
		password = env
	}
	if env, ok := os.LookupEnv("PG_DBNAME"); ok {
		dbname = env
	}
	if env, ok := os.LookupEnv("PG_PORT"); ok {
		port = env
	}
	if env, ok := os.LookupEnv("PG_SSL"); ok {
		sslmode = env
	}
	if env, ok := os.LookupEnv("PG_TIMEZONE"); ok {
		timezone = env
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error" + err.Error())
		panic(err)
	}

	return db
}
