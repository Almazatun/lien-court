package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/almazatun/lien-court/pkg/common/helper"
)

var DB *sql.DB

func ConnectDB() error {
	var err error

	// because our config function returns a string, we are parsing our str to int here - hack ??
	port, err := strconv.ParseUint(helper.GetEnvVar("DB_PORT"), 10, 32)

	if err != nil {
		return errors.New("Error parsing str to int")
	}

	host := helper.GetEnvVar("DB_HOST")
	user := helper.GetEnvVar("DB_USER")
	dbName := helper.GetEnvVar("DB_NAME")
	pass := helper.GetEnvVar("DB_PASS")

	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbName))

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Successfully connect database ✅")
	return nil
}
