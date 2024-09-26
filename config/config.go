package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	Port   = 8080
	OIDCURL = "https://auth.login.yahoo.co.jp"
)

var (
	ClientID     string
	ClientSecret string
	RedirectURI  string
)

func init() {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}

	// Initialize dynamic values
	ClientID = os.Getenv("YAHOO_CLIENT_ID")
	ClientSecret = os.Getenv("YAHOO_CLIENT_SECRET")
	RedirectURI = fmt.Sprintf("http://localhost:%d/callback", Port)
}
