package main

import (
	"github.com/FarhanRizkiM/test-be/url"
	"net/http"
)

func main() {
	http.HandleFunc("/", url.Web)
	http.ListenAndServe(":8080", nil)
}