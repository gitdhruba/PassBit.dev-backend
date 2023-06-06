package util

//This package contains function for creating jwt tokens , verifying Googleaccesstoken and middleware
//Author : Dhruba Sinha

import (
	"encoding/json"
	"io"
	"net/http"
	"passbit/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// This function returns the access-token based on the given username
func GenerateAccessToken(username string) (string, error) {

	//get the jwt-secret
	jwtsecretkey := []byte(config.Config("JWTSECRET"))

	//build claims
	t := time.Now()
	claims := jwt.StandardClaims{
		Issuer:    username,
		ExpiresAt: t.Add(30 * time.Minute).Unix(),
		Subject:   "access_token",
		IssuedAt:  t.Unix(),
	}

	//generate tokenstring
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtsecretkey)

	return tokenstring, err
}

// This function verifies the Google Idtoken
func VerifyGoogleAccessToken(accesstoken string) (string, string, bool, bool) {

	//call googleapis endpoint to verify the token
	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accesstoken)
	if err != nil {
		return "", "", false, true
	}

	defer res.Body.Close()

	//read response body
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", false, true
	}

	//parse claims
	var claims map[string]interface{}
	if err := json.Unmarshal([]byte(string(resbody)), &claims); err != nil {
		return "", "", false, true
	}

	if claims["error"] != nil {
		return "", "", false, true
	}

	return claims["name"].(string), claims["email"].(string), claims["verified_email"].(bool), false
}
