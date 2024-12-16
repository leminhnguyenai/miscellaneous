package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	parsedUrl, err := url.Parse(
		fmt.Sprintf("http://localhost%s%s", os.Getenv("PORT"), r.URL.String()),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	queryParams := parsedUrl.Query()
	code := queryParams.Get("code")

	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken := token.RefreshToken

	cookie := http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   120,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	log.Println(cookie)

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/dashboard", http.StatusPermanentRedirect)
}
