package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	HOST     = ""
	USER     = ""
	PASSWORD = ""
	DBNAME   = ""
	PORT     = 0
	TIMEZONE = ""
	DSN      = ""
)

func Load() string {
	var err error

	HOST = os.Getenv("HOST")

	USER = os.Getenv("DATABASE_USER")

	PASSWORD = os.Getenv("DATABASE_PASSWORD")

	DBNAME = os.Getenv("DATABASE_NAME")

	PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {

		panic(err)
	}
	TIMEZONE = os.Getenv("TIME_ZONE")

	DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		HOST,
		USER,
		PASSWORD,
		DBNAME,
		PORT,
		TIMEZONE,
	)

	return DSN

}
