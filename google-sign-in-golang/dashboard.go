package main

import (
	"context"
	"html/template"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jwtTokenString := cookie.Value

	jwtTokenClaims, err := VerifyToken(jwtTokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken, ok := jwtTokenClaims["refreshToken"].(string)
	if refreshToken == "" || !ok {
		http.Error(w, "Error extracting refresh token", http.StatusInternalServerError)
		return
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
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

	tokenSource := conf.TokenSource(ctx, token)

	client := oauth2.NewClient(ctx, tokenSource)

	oauth2Service, err := oauth2api.NewService(
		ctx,
		option.WithHTTPClient(client),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	email := userInfo.Email
	if email == "" {
		http.Error(
			w,
			"Can't retrieve user's email",
			http.StatusInternalServerError,
		)
	}

	templ, err := template.ParseFiles("dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Email string
	}{
		Email: email,
	}

	templ.Execute(w, data)
}
