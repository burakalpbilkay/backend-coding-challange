package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type Action struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	UserID     int    `json:"userId"`
	TargetUser int    `json:"targetUser,omitempty"`
	CreatedAt  string `json:"createdAt"`
}

var users []User
var actions []Action

func loadData() {
	userData, err := os.ReadFile("users.json")

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(userData, &users)

	if err != nil {
		log.Fatal(err)
	}

	actionData, err := os.ReadFile("actions.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(actionData, &actions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data loaded successfully.")

}

func main() {
	loadData()

	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", getUserByID).Methods("GET")
	r.HandleFunc("/user/{id}/actions/count", getUserActionCount).Methods("GET")
	r.HandleFunc("/action/{type}/next", getNextActionProbabilities).Methods("GET")
	r.HandleFunc("/users/referral-index", getReferralIndex).Methods("GET")

	http.Handle("/", r)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getUserByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	for _, user := range users {

		if fmt.Sprintf("%d", user.ID) == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func getUserActionCount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	count := 0
	for _, action := range actions {
		if fmt.Sprintf("%d", action.UserID) == id {
			count++
		}
	}

	json.NewEncoder(w).Encode(map[string]int{"count": count})
}

func getNextActionProbabilities(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	actionType := params["type"]

	nextActionCount := make(map[string]int)
	total := 0

	for i := 0; i < len(actions)-1; i++ {
		if actions[i].Type == actionType && actions[i].UserID == actions[i+1].UserID {
			nextAction := actions[i+1].Type
			nextActionCount[nextAction]++
			total++
		}
	}

	probabilities := make(map[string]float64)
	for action, count := range nextActionCount {
		probabilities[action] = float64(count) / float64(total)
	}

	json.NewEncoder(w).Encode(probabilities)
}

func getReferralIndex(w http.ResponseWriter, r *http.Request) {
	referralMap := make(map[int][]int)

	// Populate the referralMap from actions
	for _, action := range actions {
		if action.Type == "REFER_USER" {
			referralMap[action.UserID] = append(referralMap[action.UserID], action.TargetUser)
		}
	}

	referralCountCache := make(map[int]int)

	// Calculate the referral count for each user
	for _, user := range users {
		userID := user.ID
		//Skip if already calculated for the user
		if _, exists := referralCountCache[userID]; exists {
			continue
		}

		// Using BFS approach
		// Initialize a queue
		queue := []int{userID}
		visited := make(map[int]bool)
		totalReferrals := 0

		for len(queue) > 0 {
			currentUser := queue[0]
			queue = queue[1:]

			// Skip if this user has already processed
			if visited[currentUser] {
				continue
			}
			visited[currentUser] = true

			// Add direct referrals to the count and queue
			totalReferrals += len(referralMap[currentUser])
			for _, referredUser := range referralMap[currentUser] {
				if _, calculated := referralCountCache[referredUser]; !calculated {
					queue = append(queue, referredUser)
				}
			}
		}

		referralCountCache[userID] = totalReferrals
	}

	json.NewEncoder(w).Encode(referralCountCache)
}
