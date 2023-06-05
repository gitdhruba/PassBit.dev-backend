package routes

//This package contains code for setting up the endpoints or routes
//Author : Dhruba Sinha

import "github.com/gofiber/fiber/v2"

var Auth fiber.Router      //for /passbitapi/auth/*
var Protected fiber.Router //for /passbitapi/protected/*

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "hello dhruba",
		})
	})

	passbitapi := app.Group("/passbitapi")

	Auth = passbitapi.Group("/auth")
	Protected = passbitapi.Group("/protected")

	CreateEndpoints()
}
