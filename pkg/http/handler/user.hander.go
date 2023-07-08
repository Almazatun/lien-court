package handler

import (
	"log"

	inputs "github.com/almazatun/lien-court/pkg/common"
	"github.com/almazatun/lien-court/pkg/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerInstance struct{}

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
}

func (uh *UserHandlerInstance) Register(c *fiber.Ctx) error {

	// Instantiate new User struct
	u := inputs.Register{}

	//  Parse body into product struct
	if err := c.BodyParser(&u); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}

	res, err := services.CreateUser(u)

	if err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "✅",
		"res":     res,
	})

	return nil
}

func (uh *UserHandlerInstance) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := services.GetUser(id)

	if err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "✅",
		"res":     res,
	})

	return nil
}
