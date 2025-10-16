package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type APIResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func FetchCatFacts() string {
	url := "https://catfact.ninja/fact"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("error fetching cat fact:", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Print("Error: unexpected status code: ", response.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("error reading response body: ", err)
		return ""
	}

	var catFactResponse APIResponse

	err = json.Unmarshal(body, &catFactResponse)

	if err != nil {
		fmt.Println("error parsing JSON: ", err)
		return ""
	}

	return catFactResponse.Fact

}
