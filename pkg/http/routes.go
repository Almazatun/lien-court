package routes

import (
	"github.com/almazatun/lien-court/pkg/http/handler"
	middleware "github.com/almazatun/lien-court/pkg/middlware"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1/")

	// init handlers
	userHandler := handler.UserHandlerInstance{}
	authHandler := handler.AuthHandlerInstance{}
	linkHandler := handler.LinkHandlerInstance{}

	//User
	route.Post("users/register", userHandler.Register)
	// route.Get("/users/:id", userHandler.GetUser)

	// Auth
	route.Post("auth/login", authHandler.Login)
	route.Get("auth/me", middleware.JWTProtected(), authHandler.Me)

	// Link
	route.Get("links", middleware.JWTProtected(), linkHandler.List)
	route.Get("links/:id", middleware.JWTProtected(), linkHandler.GetLink)
}
