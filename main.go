package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FarhanRizkiM/test-be/controller"
)

func main() {
	// Set environment variables
	os.Setenv("MONGOCONNSTRINGENV", "mongodb+srv://Farhanrizki:NWoNivDUSqNb1Kc2@rizki.vyqexpk.mongodb.net/")
	os.Setenv("PASETOPRIVATEKEYENV", "55f8e2c264dd71fa3513a68ffd88d3060728c2c5f25bf29f99993c17ab2d5780626651fd3acdc7ae7f53c9fff589390cd2ba8b60fe2db46a90f9c6dd61bcb2ff") // Sesuaikan dengan kunci pribadi PASETO Anda

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		response := controller.Register("MONGOCONNSTRINGENV", "testSeterahBot", r) // Sesuaikan dengan nama database Anda
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		response := controller.Login("PASETOPRIVATEKEYENV", "MONGOCONNSTRINGENV", "testSeterahBot", "user", r) // Sesuaikan dengan nama database dan koleksi Anda
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
