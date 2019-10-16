package feedback

import (
	"log"
	"net/http"

	"github.com/omaressameldin/feedback-ninja/app/pkg/env"
	"github.com/omaressameldin/feedback-ninja/app/pkg/reply"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	env.ValidateEnvKeys()

	if r.Method == http.MethodPost {
		body, err := decodeBody(r)
		if err != nil {
			log.Println(err)
			return
		}

		if body.Challenge != "" {
			w.Write([]byte(body.Challenge))
			return
		}

		if isSentByBot(body.Event) {
			return
		}

		if body.Event != nil {
			err := sendFeedback(body.Event.Text)
			if err != nil {
				reply.SendErrorMessage(body.Event.Channel, err.Error())
			}

			reply.SendSuccessMessage(body.Event.Channel, messageSent)
		}
	}
}

func sendFeedback(message string) error {
	feedbackChannel := env.GetFeedbackChannelID()
	return reply.SendInfoMessage(feedbackChannel, message)
}
