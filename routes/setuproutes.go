package routes

//This package contains code for setting up the endpoints or routes
//Author : Dhruba Sinha

import (
	"passbit/util"

	"github.com/gofiber/fiber/v2"
)

var Auth fiber.Router      //router for /passbitapi/auth/*
var Protected fiber.Router //router for /passbitapi/protected/*

// setup all endpoints and groups
func SetupRoutes(app *fiber.App) {

	//don't allow any request to "/"
	app.All("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusForbidden)
	})

	//create group for api-endpoints
	passbitapi := app.Group("/passbitapi")

	//separate group for auth
	Auth = passbitapi.Group("/auth")

	//separate group for protected endpoints
	Protected = passbitapi.Group("/protected")
	Protected.Use(util.Middleware()) //apply middleware

	//release endpoints with handlers
	CreateEndpoints()
}
