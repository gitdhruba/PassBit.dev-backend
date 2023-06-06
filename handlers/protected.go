package handlers

//This package contains all handler functions for endpoints
//Author : Dhruba Sinha

import (
	"fmt"
	db "passbit/database"
	"passbit/models"

	"github.com/gofiber/fiber/v2"
)

// handler for "/passbitapi/protected/setmasterpasswd"
func SetMasterpasswd(c *fiber.Ctx) error {

	//define struct for request body
	type RequestBody models.Mpass

	var req RequestBody
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("ERROR : ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	//check wheather the user exists or not
	if dbres := db.DB.Where(&models.User{Username: req.Username}).First(&models.User{}); dbres.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "user doesn't exists",
		})
	}

	//check if master-password has already been set
	var mpass models.Mpass
	mpass.Username = req.Username
	if dbres := db.DB.Where(&mpass).First(&mpass); dbres.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "master password already set",
		})
	}

	//set master password
	mpass.Masterpasswd = req.Masterpasswd
	if err := db.DB.Create(&mpass).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "master password set successfully",
	})
}
