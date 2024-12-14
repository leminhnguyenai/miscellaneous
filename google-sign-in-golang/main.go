package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load(
		"/Users/minhnl2012/Documents/Projects/miscellaneous/google-sign-in-golang/.env",
	)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})

	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/callback", Callback)
	mux.HandleFunc("/dashboard", Dashboard)

	port := os.Getenv("PORT")

	log.Printf("The server is on http://localhost%s\n", port)
	http.ListenAndServe(port, mux)
}
