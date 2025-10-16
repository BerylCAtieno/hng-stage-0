package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BerylCatieno/me/data"
)

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Stack string `json:"stack"`
}

var ResponsePayload struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	Timestamp string `json:"timestamp"`
	Fact      string `json:"fact"`
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {

	catFact := data.FetchCatFacts()
	now := time.Now().UTC()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed: ", http.StatusMethodNotAllowed)
		return
	}

	var user User

	user.Email = "berylatieno30@gmail.com"
	user.Name = "Beryl Atieno"
	user.Stack = "Go/Gin"

	ResponsePayload.Status = "success"
	ResponsePayload.User = user
	ResponsePayload.Timestamp = now.String()
	ResponsePayload.Fact = catFact

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ResponsePayload)
}
