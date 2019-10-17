package reply

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/feedback-ninja/app/pkg/env"
)

type Reply struct {
	Attachments []slack.Attachment `json:"attachments"`
	Blocks      []slack.Block      `json:"blocks"`
}

func sendReply(channelID string, r Reply) error {
	token := env.GetToken()
	api := slack.New(token)
	var err error
	options := []slack.MsgOption{
		slack.MsgOptionAttachments(r.Attachments...),
		slack.MsgOptionBlocks(r.Blocks...),
	}
	_, _, err = api.PostMessage(channelID, options...)

	return err
}

func replaceMessage(
	responseURL string,
	w http.ResponseWriter,
	attachments []slack.Attachment,
) error {
	jsonValue, _ := json.Marshal(Reply{
		Attachments: attachments,
	})

	_, err := http.Post(responseURL, "application/json", bytes.NewBuffer(jsonValue))
	return err
}

func SendErrorMessage(channelID string, text string) {
	err := sendReply(
		channelID,
		Reply{Attachments: []slack.Attachment{{Text: text, Color: colorDanger}}},
	)

	if err != nil {
		log.Println(err)
	}
}

func SendActionError(responseURL string, w http.ResponseWriter, text string) error {
	return replaceMessage(
		responseURL,
		w,
		[]slack.Attachment{{Text: text, Color: colorDanger}},
	)
}

func SendActionSuccess(responseURL string, w http.ResponseWriter, text string) error {
	return replaceMessage(
		responseURL,
		w,
		[]slack.Attachment{{Text: text, Color: colorSuccess}},
	)
}

func SendInfoMessage(channelID string, text string) error {
	return sendReply(
		channelID,
		Reply{Attachments: []slack.Attachment{{Text: text, Color: colorInfo}}},
	)
}

func SendActions(channelID string, blocks []slack.Block) error {
	return sendReply(
		channelID,
		Reply{Blocks: blocks},
	)
}
