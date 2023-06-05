package handlers

//This package contains all handler functions for endpoints
//Author : Dhruba Sinha

import (
	"fmt"
	db "passbit/database"
	"passbit/models"
	"passbit/util"

	"github.com/gofiber/fiber/v2"
)

// handler for "/passbitapi/auth/signin"
func SigninUser(c *fiber.Ctx) error {

	//define struct for signin request body
	type RequestBody struct {
		User              string `json:"username"`
		Email             string `json:"email"`
		Googleaccesstoken string `json:"googleaccesstoken"`
	}

	req := new(RequestBody)
	if err := c.BodyParser(req); err != nil {
		fmt.Printf("ERROR1 : %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	//verify Googleaccesstoken
	userfromtoken, emailfromtoken, isverifiedemail, err := util.VerifyGoogleAccessToken(req.Googleaccesstoken)
	if err != nil {
		fmt.Printf("ERROR2 : %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	//check wheather the email is verified or not
	if !isverifiedemail {
		fmt.Println("ERROR : email is not verified")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "email is not verified",
		})
	}

	//match user and email
	if req.User != userfromtoken || req.Email != emailfromtoken {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "user details provided is not compatible with idtoken",
		})
	}

	//make db entry for the user
	var user models.User
	user.Username = userfromtoken
	user.Email = emailfromtoken
	if dbres := db.DB.Where(&user).First(&user); dbres.RowsAffected <= 0 {
		err := db.DB.Create(&user).Error
		if err != nil {
			fmt.Printf("DBERROR : %s", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err,
			})
		}
	}

	//generate access-token
	accesstoken, err := util.GenerateAccessToken(userfromtoken)
	if err != nil {
		fmt.Printf("ERROR3 : %s", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	//fmt.Println(accesstoken)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":       false,
		"accesstoken": accesstoken,
	})

}
