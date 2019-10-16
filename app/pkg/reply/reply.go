package reply

import (
	"log"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/feedback-ninja/app/pkg/env"
)

type Reply struct {
	Attachments []slack.Attachment `json:"attachments"`
}

func sendReply(channelID string, r Reply) error {
	token := env.GetToken()
	api := slack.New(token)

	_, _, err := api.PostMessage(channelID, slack.MsgOptionAttachments(
		r.Attachments...,
	))

	return err
}

func SendErrorMessage(channelID string, text string) {
	err := sendReply(channelID, Reply{
		Attachments: []slack.Attachment{
			{
				Text:  text,
				Color: colorDanger,
			},
		},
	})

	if err != nil {
		log.Println(err)
	}
}

func SendSuccessMessage(channelID string, text string) error {
	return sendReply(channelID, Reply{
		Attachments: []slack.Attachment{
			{
				Text:  text,
				Color: colorSuccess,
			},
		},
	})
}

func SendInfoMessage(channelID string, text string) error {
	return sendReply(channelID, Reply{
		Attachments: []slack.Attachment{
			{
				Text:  text,
				Color: colorInfo,
			},
		},
	})
}
