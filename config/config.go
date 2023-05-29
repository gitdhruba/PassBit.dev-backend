package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Loadenv() {

	//load env file
	err := godotenv.Load("./app.env")
	if err != nil {
		fmt.Println("ERROR: could not load env file")
		os.Exit(1)
	}
}

func Config(key string) string {

	return os.Getenv(key)

}
