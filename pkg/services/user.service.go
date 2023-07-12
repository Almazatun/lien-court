package services

import (
	"errors"
	"fmt"
	"log"

	inputs "github.com/almazatun/lien-court/pkg/common"
	"github.com/almazatun/lien-court/pkg/common/helper"
	"github.com/almazatun/lien-court/pkg/database"
	"github.com/almazatun/lien-court/pkg/database/entities"
)

func CreateUser(registerInput inputs.Register) (r *inputs.Register, err error) {
	u := entities.User{}
	fmt.Println(registerInput)

	row, err := database.DB.Query(`SELECT * FROM "user" WHERE email = $1`, registerInput.Email)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Pass, &u.CreatedAt)

		if err != nil {
			log.Println("No rows were returned!")

			return nil, err
		}
	}

	if u.Email == registerInput.Email {
		return nil, errors.New(`User already exists by email`)
	}

	passHash, err := helper.GenPassHash(registerInput.Pass)

	if err != nil {
		return nil, err
	}

	// Insert User into database
	res, err := database.DB.Query(
		`INSERT INTO "user" (username, email, pass) VALUES ($1, $2, $3)`,
		registerInput.Name, registerInput.Email, passHash,
	)

	if err != nil {
		return nil, err
	}

	// check result
	log.Println(res)

	return &registerInput, nil
}

func GetUser(id string) (res *entities.User, err error) {
	u := entities.User{}

	row, err := database.DB.Query(`SELECT * FROM "user" WHERE id = $1`, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Pass, &u.CreatedAt)

		if err != nil {
			log.Println("No rows were returned!")

			return nil, err
		}
	}

	// check result
	log.Println(u)

	return &u, nil
}
