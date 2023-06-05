// REST API for PassBit project
//Author : Dhruba Sinha

package main

//import required packages
import (
	"fmt"
	"os"
	"passbit/config"
	"passbit/database"
	"passbit/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// main entrypoint
func main() {

	//Load env file
	config.Loadenv()

	//connect to db and automigrate all models
	database.ConnectDB()
	database.AutomigrateModels()

	//create new fiber app instance
	app := fiber.New()
	//apply default CORS middleware for now
	app.Use(cors.New())

	//setup routes
	routes.SetupRoutes(app)

	//start listening to PORT
	portstring := fmt.Sprintf(":%s", config.Config("PORT"))
	if err := app.Listen(portstring); err != nil {
		fmt.Println("ERROR : could not start the server")
		os.Exit(1)
	}

}
