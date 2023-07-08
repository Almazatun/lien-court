package services

import (
	"errors"
	"log"

	inputs "github.com/almazatun/lien-court/pkg/common"
	"github.com/almazatun/lien-court/pkg/common/helper"
	"github.com/almazatun/lien-court/pkg/database"
	"github.com/almazatun/lien-court/pkg/database/entities"
)

func Login(loginInput inputs.Login) (r *entities.User, err error) {
	u := entities.User{}

	row, err := database.DB.Query(`SELECT * FROM "user" WHERE email = $1`, loginInput.Email)

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

	// check password
	isCorrectPass := helper.ComparePasswords(u.Pass, loginInput.Pass)

	if !isCorrectPass {
		return nil, errors.New("Incorrect password")
	}

	// check result
	log.Println(u)

	return &u, nil
}
