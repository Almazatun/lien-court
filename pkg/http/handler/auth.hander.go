package handler

import (
	"fmt"
	"log"
	"strings"

	inputs "github.com/almazatun/lien-court/pkg/common"
	"github.com/almazatun/lien-court/pkg/common/helper"
	jwt_ "github.com/almazatun/lien-court/pkg/common/jwt"
	"github.com/almazatun/lien-court/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthHandlerInstance struct{}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Me(c *fiber.Ctx) error
}

func (ah *AuthHandlerInstance) Login(c *fiber.Ctx) error {
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

	// Validation
	errors, err := helper.ValidateInputModel[inputs.Login](loginUserInput)

	if err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"errors":  errors,
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

func (ah *AuthHandlerInstance) Me(c *fiber.Ctx) error {

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

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": claims,
	})

	return nil
}
