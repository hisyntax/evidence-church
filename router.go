package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func initialiseRouter() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/sermon", sermonHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
