package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PushEvent struct {
	Ref     string `json:"ref"`
	Before  string `json:"before"`
	After   string `json:"after"`
	Created bool   `json:"created"`
	Deleted bool   `json:"deleted"`
	Forced  bool   `json:"forced"`
	Compare string `json:"compare"`
	Commits []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
		Author  struct {
			Name string `json:"name"`
		} `json:"author"`
	} `json:"commits"`
	HeadCommit *struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"head_commit"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Pusher struct {
		Name string `json:"name"`
	} `json:"pusher"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Respond quickly with 202 Accepted
	w.WriteHeader(http.StatusAccepted)

	// Get GitHub event type
	githubEvent := r.Header.Get("X-GitHub-Event")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		return
	}
	defer r.Body.Close()

	switch githubEvent {
	case "push":
		var event PushEvent
		if err := json.Unmarshal(body, &event); err != nil {
			log.Println("Error parsing JSON:", err)
			return
		}
		fmt.Printf("Got push event: %s, %s\n", githubEvent, string(body))
	default:
		fmt.Printf("Unhandled event: %s, %s\n", githubEvent, string(body))
	}
}
