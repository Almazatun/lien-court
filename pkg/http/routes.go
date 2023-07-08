package routes

import (
	"github.com/almazatun/lien-court/pkg/http/handler"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// init handlers
	userHandler := handler.UserHandlerInstance{}
	authHandler := handler.AuthHandlerInstance{}

	//User
	route.Post("/users/register", userHandler.Register)
	// route.Get("/users/:id", userHandler.Get)

	// Auth
	route.Post("/auth/login", authHandler.Login)
}
