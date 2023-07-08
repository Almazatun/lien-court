package handler

import (
	"fmt"
	"log"

	inputs "github.com/almazatun/lien-court/pkg/common"
	jwt_ "github.com/almazatun/lien-court/pkg/common/jwt"
	"github.com/almazatun/lien-court/pkg/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandlerInstance struct{}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

func (uh *AuthHandlerInstance) Login(c *fiber.Ctx) error {
	loginUserInput := inputs.Login{}

	//  Parse body into product struct
	if err := c.BodyParser(&loginUserInput); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}

	res, err := services.Login(loginUserInput)

	if err != nil {
		log.Println(err)

		if res == nil {
			c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
			return nil
		}

		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}

	id := fmt.Sprintf("%v", res.ID)

	t, err := jwt_.GenToken(res.Email, id)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "✅",
		"res": fiber.Map{
			"accessToken": t,
		},
	})

	return nil
}
