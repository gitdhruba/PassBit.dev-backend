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

	var req RequestBody
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("ERROR : ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	//verify Googleaccesstoken
	userfromtoken, emailfromtoken, isverifiedemail, err := util.VerifyGoogleAccessToken(req.Googleaccesstoken)
	if err != nil {
		fmt.Println("ERROR : ", err)
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
			fmt.Println("DBERROR : ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err,
			})
		}
	}

	//generate access-token
	accesstoken, err := util.GenerateAccessToken(userfromtoken)
	if err != nil {
		fmt.Println("ERROR : ", err)
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

// handler for "/passbitapi/auth/reissueaccesstoken"
func ReIssueAccesstoken(c *fiber.Ctx) error {

	//define struct for request body
	type RequestBody struct {
		Googleaccesstoken string `json:"googleaccesstoken"`
	}

	var req RequestBody
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("ERROR : ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	//verify googleaccesstoken
	username, email, isverifiedemail, err := util.VerifyGoogleAccessToken(req.Googleaccesstoken)
	if err != nil {
		fmt.Println("ERROR : ", err)
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

	//check if the user is registered or not
	var user models.User
	user.Username = username
	user.Email = email
	if dbres := db.DB.Where(&user).First(&user); dbres.RowsAffected <= 0 {
		//not registered
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "user is not registered",
		})
	}

	//regenerate accesstoken
	accesstoken, err := util.GenerateAccessToken(username)
	if err != nil {
		fmt.Println("ERROR : ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":       false,
		"accesstoken": accesstoken,
	})
}
