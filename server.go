package main

//import required packages
import (
	"fmt"
	"passbit/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// main entrypoint
func main() {

	//create new fiber app instance
	app := fiber.New()

	//apply default CORS middleware for now
	app.Use(cors.New())

	//start listening to PORT
	portstring := fmt.Sprintf(":%s", config.Config("PORT"))
	fmt.Printf("Listening on port %s", portstring)
	app.Listen(portstring)

	return
}
