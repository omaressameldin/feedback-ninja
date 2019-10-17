package confirmation

import (
	"log"
	"net/http"

	"github.com/omaressameldin/feedback-ninja/app/pkg/env"
	"github.com/omaressameldin/feedback-ninja/app/pkg/feedback"
	"github.com/omaressameldin/feedback-ninja/app/pkg/reply"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	env.ValidateEnvKeys()

	if r.Method == http.MethodPost {
		payload, _ := unmarshalPayload(r)
		if mightCancel(w, payload) {
			return
		}

		for _, block := range payload.ActionCallback.BlockActions {
			switch block.BlockID {
			case feedback.FeedbackBlockID:
				{
					sendFeedback(block.Value, payload.ResponseURL, w)
				}
			}
		}
	}
}

func sendFeedback(message string, responseUrl string, w http.ResponseWriter) {
	feedbackChannel := env.GetFeedbackChannelID()
	if err := reply.SendInfoMessage(feedbackChannel, message); err != nil {
		reply.SendActionError(responseUrl, w, err.Error())
	}

	if err := reply.SendActionSuccess(responseUrl, w, messageSent); err != nil {
		log.Println(err)
	}
}
