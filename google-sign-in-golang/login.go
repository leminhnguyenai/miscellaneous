package main

import (
	"net/http"
	"os"

	"github.com/lpernett/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Login(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(
		"/Users/minhnl2012/Documents/Projects/miscellaneous/google-sign-in-golang/.env",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	googleConsentURL := conf.AuthCodeURL(
		"stake-token",
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)

	http.Redirect(w, r, googleConsentURL, http.StatusTemporaryRedirect)
}
