package services

import (
	"errors"
	"log"

	inputs "github.com/almazatun/lien-court/pkg/common"
	"github.com/almazatun/lien-court/pkg/database"
	"github.com/almazatun/lien-court/pkg/database/entities"
)

func LinkList(input inputs.LinkList, userId string) (r *entities.Links, err error) {
	result := entities.Links{Links: []entities.Link{}}

	row, err := database.DB.Query(
		`SELECT * FROM link WHERE user_id = $1 OFFSET $2 LIMIT $3`,
		userId, input.Limit, input.Page,
	)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		link := entities.Link{}
		err := row.Scan(&link.ID, &link.Original, &link.Short, &link.CreatedAt)

		if err != nil {
			log.Println("No rows were returned!")

			return nil, err
		}

		result.Links = append(result.Links, link)
	}

	log.Println(result.Links)

	return &result, nil
}

func GetLink(id string) (r *entities.Link, err error) {
	link := entities.Link{}

	row, err := database.DB.Query(`SELECT * FROM link WHERE id = $2`, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		err := row.Scan(&link.ID, &link.Original, &link.Short, &link.CreatedAt)

		if err != nil {
			log.Println("No rows were returned!")

			return nil, err
		}
	}

	if link.ID.String() != id {
		return nil, errors.New("Invalid id")
	}

	log.Println(link)

	return &link, nil
}
