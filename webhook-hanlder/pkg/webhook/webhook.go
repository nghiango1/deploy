package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nghiango1/deploy/webhook-handler/pkg/logger"
	"github.com/nghiango1/deploy/webhook-handler/pkg/webhook/push"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Respond quickly with 202 Accepted
	w.WriteHeader(http.StatusAccepted)

	// Get GitHub event type
	githubEvent := r.Header.Get("X-GitHub-Event")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Get().Debug("Error reading request body:", err)
		return
	}
	defer r.Body.Close()

	switch githubEvent {
	case "push":
		var event push.Event
		if err := json.Unmarshal(body, &event); err != nil {
			logger.Get().Debug("Error parsing JSON:", err)
			return
		}
		switch event.Ref {
		case "refs/heads/main":
			logger.Get().Debug("Main ref push event\n")
			push.MainRefHandler(event)
		case "refs/heads/dev":
			logger.Get().Debug("Dev ref push event\n")
			push.DevRefHandler(event)
		default:
			logger.Get().Warn(fmt.Sprintf("Unhandled ref: %s\n", event.Ref))
		}
	default:
		logger.Get().Warn(fmt.Sprintf("Unhandled event: %s, %s\n", githubEvent))
	}
}
