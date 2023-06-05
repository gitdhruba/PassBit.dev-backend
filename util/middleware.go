package util

//This package contains function for creating jwt tokens , verifying Googleaccesstoken and middleware
//Author : Dhruba Sinha

import (
	"passbit/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Middleware returns middleware for protected routes
func Middleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		//get accesstoken from Headers
		accesstoken := c.GetReqHeaders()["access_token"]

		//parse claims
		claims := jwt.StandardClaims{}
		token, _ := jwt.ParseWithClaims(accesstoken, &claims,
			func(token *jwt.Token) (interface{}, error) {
				return config.Config("JWTSECRET"), nil
			})

		//check for validation
		if token.Valid {
			if claims.ExpiresAt < time.Now().Unix() {
				//token expired
				return c.SendStatus(fiber.StatusUnauthorized)
			}
		} else {
			//token is not valid
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		//validation done , allow request
		return c.Next()
	}
}
