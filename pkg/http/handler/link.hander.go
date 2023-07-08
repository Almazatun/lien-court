package handler

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	inputs "github.com/almazatun/lien-court/pkg/common"
	"github.com/almazatun/lien-court/pkg/common/helper"
	"github.com/almazatun/lien-court/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type LinkHandlerInstance struct{}

type LinkHandler interface {
	List(c *fiber.Ctx) error
	GetLink(c *fiber.Ctx) error
}

func (lh *LinkHandlerInstance) List(c *fiber.Ctx) error {
	input := inputs.LinkList{}

	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"message": "You are not logged in",
		})
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(helper.GetEnvVar("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"message": "invalid token claim",
		})

	}

	// log.Println(claims)

	userId := fmt.Sprintf("%v", claims["id"])

	//  Parse body into product struct
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}

	if reflect.ValueOf(input).IsZero() {
		input.Limit = 10
	}

	if reflect.ValueOf(input).IsZero() {
		input.Page = (1 - 1) * input.Limit
	} else {
		input.Page = (input.Page - 1) * input.Limit
	}

	res, err := services.LinkList(input, userId)

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

func (lh *LinkHandlerInstance) GetLink(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := services.GetLink(id)

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
