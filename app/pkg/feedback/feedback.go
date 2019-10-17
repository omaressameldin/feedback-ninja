package feedback

import (
	"log"
	"net/http"

	"github.com/nlopes/slack"
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

		if !isSentByUser(body.Event) {
			return
		}

		if body.Event != nil {
			err := sendConfirmation(body.Event.Channel, body.Event.Text, nil)
			if err != nil {
				reply.SendErrorMessage(body.Event.Channel, err.Error())
			}
		}
	}
}

func sendConfirmation(channelID string, message string, time *string) error {
	return reply.SendActions(channelID,
		[]slack.Block{
			slack.NewContextBlock(
				"",
				[]slack.MixedElement{
					slack.NewTextBlockObject("mrkdwn", confirmationMessage, false, false),
				}...,
			),
			slack.NewActionBlock(
				FeedbackBlockID,
				reply.PrimaryButton(confirmButton, message),
				reply.CancelButton(),
			),
		},
	)
}
