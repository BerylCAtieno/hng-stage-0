package main

import (
	"fmt"
	"net/http"

	"github.com/BerylCatieno/me/api"
)

func main() {
	http.HandleFunc("/me", api.GetProfileHandler)

	url := "localhost"
	port := "8080"

	fmt.Printf("server running on: %s, port: %s", url, port)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
