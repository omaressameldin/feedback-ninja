package confirmation

import (
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/feedback-ninja/app/pkg/reply"
)

func unmarshalPayload(r *http.Request) (*slack.InteractionCallback, error) {
	var payload slack.InteractionCallback
	err := json.Unmarshal([]byte(r.FormValue("payload")), &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func mightCancel(w http.ResponseWriter, payload *slack.InteractionCallback) bool {
	isCancel := payload.ActionCallback.BlockActions[0].Value == reply.CancelValue
	if isCancel {
		reply.SendActionError(
			payload.ResponseURL,
			w,
			reply.CancelMessage,
		)
	}

	return isCancel
}
