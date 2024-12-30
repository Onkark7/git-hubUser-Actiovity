package opra

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Event struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
}

func FetchUserActivity(username string) {

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching GitHub activity: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: GitHub API returned status %d", resp.StatusCode)
	}

	var events []Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	fmt.Println("Recent GitHub Activity:")
	for _, event := range events {
		fmt.Printf("Event: %s, Repo: %s, Date: %s\n", event.Type, event.Repo.Name, event.CreatedAt)
	}

}
