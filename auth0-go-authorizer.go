package main

import (
	"fmt"

	auth "github.com/auth0-community/go-auth0"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func main() {
	client := auth.NewJWKClient(auth.JWKClientOptions{URI: "https://yourdomain.au.auth0.com/.well-known/jwks.json"}, nil)
	audience := "AUDIENCE"
	configuration := auth.NewConfiguration(client, []string{audience}, "https://yourdomain.au.auth0.com/", jose.RS256)
	validator := auth.NewValidator(configuration, nil)

	token, err := jwt.ParseSigned("JWT_TOKEN_COMES_HERE")
	if err != nil {
		fmt.Print("Couldn't Parse Token Err:", err)
	}

	if err := validator.ValidateToken(token); err != nil {
		fmt.Println("Cannot validate token because of", err)
	}
}
