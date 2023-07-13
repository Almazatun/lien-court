package services

import (
	"errors"
	"fmt"
	"log"

	inputs "github.com/almazatun/lien-court/pkg/common"
	"github.com/almazatun/lien-court/pkg/common/helper"
	"github.com/almazatun/lien-court/pkg/database"
	"github.com/almazatun/lien-court/pkg/database/entities"
	"github.com/google/uuid"
)

func LinkList(input inputs.LinkList, userId string) (r *entities.Links, err error) {
	result := entities.Links{Links: []entities.Link{}}
	fmt.Println("LIST", userId)

	row, err := database.DB.Query(
		`SELECT * FROM link WHERE user_id = $1 OFFSET $2 LIMIT $3`,
		userId, input.Page, input.Limit,
	)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		l := entities.Link{}
		err := row.Scan(&l.ID, &l.Original, &l.Short, &l.CreatedAt, &l.UserID)

		if err != nil {
			log.Println("No rows were returned!")

			return nil, err
		}

		result.Links = append(result.Links, l)
	}

	log.Println(result.Links)

	return &result, nil
}

func GetLink(id string) (r *entities.Link, err error) {
	l := entities.Link{}

	row, err := database.DB.Query(`SELECT * FROM link WHERE id = $1`, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		err := row.Scan(&l.ID, &l.Original, &l.Short, &l.CreatedAt, &l.UserID)

		if err != nil {
			log.Println("No rows were returned!")

			return nil, err
		}
	}

	if l.ID.String() != id {
		return nil, errors.New("Invalid id")
	}

	log.Println(l)

	return &l, nil
}

func CreateLink(userId, link string) (r *entities.Link, err error) {
	l := entities.Link{}
	id := uuid.New()
	short := helper.GetEnvVar("BASE_URL_BE") + "/api/v1/links/" + id.String()

	res, err := database.DB.Query(
		`INSERT INTO "link" (id, short, original, user_id) VALUES ($1, $2, $3, $4)`,
		id, link, short, userId,
	)

	defer res.Close()

	// check result
	log.Println(res)

	if err != nil {
		return nil, err
	}

	row, err := database.DB.Query(`SELECT * FROM link WHERE id = $1`, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		err := row.Scan(&l.ID, &l.Original, &l.Short, &l.CreatedAt, &l.UserID)

		if err != nil {
			log.Println("No rows were returned!")

			return nil, err
		}
	}

	// iterate through the values of the row
	for row.Next() {
		err := row.Scan(&l.ID, &l.Short, &l.Original, &l.CreatedAt, &l.UserID)

		if err != nil {
			log.Println("No rows were returned!")

			return nil, err
		}
	}

	// check result
	log.Println("Create link", l)

	return &l, nil
}
