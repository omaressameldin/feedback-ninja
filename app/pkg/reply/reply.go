package reply

import (
	"log"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/feedback-ninja/app/pkg/env"
)

type Reply struct {
	Attachments []slack.Attachment `json:"attachments"`
	Blocks      []slack.Block      `json:"blocks"`
}

func sendReply(channelID string, r Reply, time *string) error {
	token := env.GetToken()
	api := slack.New(token)
	var err error
	options := []slack.MsgOption{
		slack.MsgOptionAttachments(r.Attachments...),
		slack.MsgOptionBlocks(r.Blocks...),
	}

	if time != nil {
		_, _, _, err = api.UpdateMessage(channelID, *time, options...)
	} else {
		_, _, err = api.PostMessage(channelID, options...)
	}

	return err
}

func SendErrorMessage(channelID string, text string, time *string) {
	err := sendReply(
		channelID,
		Reply{Attachments: []slack.Attachment{{Text: text, Color: colorDanger}}},
		time,
	)

	if err != nil {
		log.Println(err)
	}
}

func SendSuccessMessage(channelID string, text string, time *string) error {
	return sendReply(
		channelID,
		Reply{Attachments: []slack.Attachment{{Text: text, Color: colorSuccess}}},
		time,
	)
}

func SendInfoMessage(channelID string, text string, time *string) error {
	return sendReply(
		channelID,
		Reply{Attachments: []slack.Attachment{{Text: text, Color: colorInfo}}},
		time,
	)
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
