package feedback

import (
	"encoding/json"
	"net/http"
)

type Body struct {
	Challenge string
	Event     *Event
}

type Event struct {
	Text    string
	Channel string
	user    string
	BotID   string `json:"bot_id"`
}

func decodeBody(r *http.Request) (*Body, error) {
	decoder := json.NewDecoder(r.Body)
	var body Body
	err := decoder.Decode(&body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func isSentByBot(e *Event) bool {
	return e != nil && e.BotID != ""
}
