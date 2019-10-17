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
	Text      string
	Channel   string
	User      string
	BotID     string `json:"bot_id"`
	Timestamp string `json:"event_ts"`
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

func isSentByUser(e *Event) bool {
	return e != nil && e.BotID == "" && e.User != ""
}
