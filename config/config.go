package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {

	//load env file
	err := godotenv.Load("./app.env")
	if err != nil {
		fmt.Println("ERROR: could not load env file")
		os.Exit(1)
	}

	return os.Getenv(key)
}
