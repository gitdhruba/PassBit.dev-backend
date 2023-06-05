package routes

//This package contains code for setting up the endpoints or routes
//Author : Dhruba Sinha

import "passbit/handlers"

func CreateEndpoints() {

	//authhandlers
	Auth.Post("/signin", handlers.SigninUser) //endpoint: "/passbitapi/auth/signin"

	//protected

}
