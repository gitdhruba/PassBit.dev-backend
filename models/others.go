package models

//This package contains all model definitions
//Author : Dhruba Sinha

//this file specifically all other model definitions except the user model

// model definition for Masterpassword
type Mpass struct {
	Username     string `json:"username" gorm:"unique"`
	Masterpasswd string `json:"master_passwd"`
}
